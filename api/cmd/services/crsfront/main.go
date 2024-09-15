package main

import (
	"context"
	"embed"
	"errors"
	"expvar"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/David-Kalashir/crs-front/api/cmd/services/crsfront/build/all"
	"github.com/David-Kalashir/crs-front/api/sdk/http/debug"
	"github.com/David-Kalashir/crs-front/api/sdk/http/mux"
	"github.com/David-Kalashir/crs-front/business/sdk/sqldb"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/minio"
	"github.com/David-Kalashir/crs-front/foundation/otel"
	"github.com/ardanlabs/conf/v3"
)

/*
	Need to figure out timeouts for http service.
*/

//go:embed all:static
var static embed.FS

var build = "develop"

//var routes = "all" // go build -ldflags "-X main.routes=crud"

func main() {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT *******")
		},
	}

	traceIDFn := func(ctx context.Context) string {
		return otel.GetTraceID(ctx)
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "CRSFRONT", traceIDFn, events)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// -------------------------------------------------------------------------
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout        time.Duration `conf:"default:5s"`
			WriteTimeout       time.Duration `conf:"default:10s"`
			IdleTimeout        time.Duration `conf:"default:120s"`
			ShutdownTimeout    time.Duration `conf:"default:20s"`
			APIHost            string        `conf:"default:0.0.0.0:7000"`
			DebugHost          string        `conf:"default:0.0.0.0:7010"`
			CORSAllowedOrigins []string      `conf:"default:*"`
		}
		DB struct {
			User         string `conf:"default:postgres"`
			Password     string `conf:"default:postgres,mask"`
			Host         string `conf:"default:database-service"`
			Name         string `conf:"default:postgres"`
			MaxIdleConns int    `conf:"default:0"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
		Tempo struct {
			Host        string  `conf:"default:tempo:4317"`
			ServiceName string  `conf:"default:crsfront"`
			Probability float64 `conf:"default:0.05"`
			// Shouldn't use a high Probability value in non-developer systems.
			// 0.05 should be enough for most systems. Some might want to have
			// this even lower.
		}
		Minio struct {
			Endpoint        string `conf:"default:localhost:9000"`
			AccessKeyID     string `conf:"default:3dbKvkVvuJxkx2xbHtEF"`
			SecretAccessKey string `conf:"default:eSmLxHOPMXeXxfcMyjp4HBtqJKdqP0z4SQYe2aMz"`
			UseSSL          bool   `conf:"default:false"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "crsfront",
		},
	}

	const prefix = "CRSFRONT"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// App Starting

	log.Info(ctx, "starting service", "version", cfg.Build)
	defer log.Info(ctx, "shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info(ctx, "startup", "config", out)

	log.BuildInfo(ctx)

	expvar.NewString("build").Set(cfg.Build)

	// -------------------------------------------------------------------------
	// Database Support

	log.Info(ctx, "startup", "status", "initializing database support", "hostport", cfg.DB.Host)

	db, err := sqldb.Open(sqldb.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}

	defer db.Close()

	// -------------------------------------------------------------------------
	// Start Tracing Support

	log.Info(ctx, "startup", "status", "initializing tracing support")

	traceProvider, err := otel.InitTracing(otel.Config{
		ServiceName: cfg.Tempo.ServiceName,
		Host:        cfg.Tempo.Host,
		ExcludedRoutes: map[string]struct{}{
			"/v1/liveness":  {},
			"/v1/readiness": {},
		},
		Probability: cfg.Tempo.Probability,
	})
	if err != nil {
		return fmt.Errorf("starting tracing: %w", err)
	}

	defer traceProvider.Shutdown(context.Background())

	tracer := traceProvider.Tracer(cfg.Tempo.ServiceName)

	// -------------------------------------------------------------------------
	// Start Debug Service

	go func() {
		log.Info(ctx, "startup", "status", "debug v1 router started", "host", cfg.Web.DebugHost)

		if err := http.ListenAndServe(cfg.Web.DebugHost, debug.Mux()); err != nil {
			log.Error(ctx, "shutdown", "status", "debug v1 router closed", "host", cfg.Web.DebugHost, "msg", err)
		}
	}()

	// -------------------------------------------------------------------------
	// create minio

	mio, err := minio.New(cfg.Minio.Endpoint, cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, cfg.Minio.UseSSL)
	if err != nil {
		return fmt.Errorf("creat minio: %w", err)
	}

	// -------------------------------------------------------------------------
	// Start API Service

	log.Info(ctx, "startup", "status", "initializing V1 API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	cfgMux := mux.Config{
		Build:  build,
		Log:    log,
		DB:     db,
		Tracer: tracer,
		Minio:  mio,
	}

	webAPI := mux.WebAPI(cfgMux,
		buildRoutes(),
		mux.WithCORS(cfg.Web.CORSAllowedOrigins),
		mux.WithFileServer(static, "static"),
	)

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      webAPI,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     logger.NewStdLogger(log, logger.LevelError),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info(ctx, "startup", "status", "api router started", "host", api.Addr)

		serverErrors <- api.ListenAndServe()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info(ctx, "shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}

func buildRoutes() mux.RouteAdder {

	// The idea here is that we can build different versions of the binary
	// with different sets of exposed web APIs. By default we build a single
	// instance with all the web APIs.
	//
	// Here is the scenario. It would be nice to build two binaries, one for the
	// transactional APIs (CRUD) and one for the reporting APIs. This would allow
	// the system to run two instances of the database. One instance tuned for the
	// transactional database calls and the other tuned for the reporting calls.
	// Tuning meaning indexing and memory requirements. The two databases can be
	// kept in sync with replication.

	return all.Routes()
}

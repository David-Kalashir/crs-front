package checkapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/checkapp"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	DB    *sqlx.DB
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	api := newAPI(checkapp.NewApp(cfg.Build, cfg.Log, cfg.DB))
	app.HandlerFuncNoMid(http.MethodGet, version, "/readiness", api.readiness)
	app.HandlerFuncNoMid(http.MethodGet, version, "/liveness", api.liveness)
}

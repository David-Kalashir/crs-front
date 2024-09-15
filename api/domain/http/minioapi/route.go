package minioapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
	"github.com/minio/minio-go/v7"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log         *logger.Logger
	MinioClient *minio.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	api := newAPI(cfg.MinioClient)
	app.RawHandlerFunc(http.MethodGet, version, "/minio/{bucketname}/{objectname}", api.getobject)

}

package uisearchpageapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/uisearchpageapp"
	"github.com/David-Kalashir/crs-front/business/domain/uisearchpagebus"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log             *logger.Logger
	UIsearchpagebus *uisearchpagebus.Business
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	api := newAPI(uisearchpageapp.NewApp(cfg.UIsearchpagebus))
	app.RawHandlerFunc(http.MethodGet, version, "/searchpage", api.searchpage)

}

package uisearchpageapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/business/domain/uisearchpagebus"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log             *logger.Logger
	UIsearchpagebus *uisearchpagebus.Business
	StoreKey        string
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	api := newAPI(cfg)
	app.RawHandlerFunc(http.MethodGet, version, "/searchpage", api.searchpage)
	app.RawHandlerFunc(http.MethodGet, version, "/searchpage/login", api.login)
	app.RawHandlerFunc(http.MethodGet, version, "/searchpage/logout", api.logout)

}

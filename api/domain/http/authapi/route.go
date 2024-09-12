package authapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/api/sdk/http/mid"
	"github.com/David-Kalashir/crs-front/app/sdk/auth"
	"github.com/David-Kalashir/crs-front/business/domain/userbus"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	UserBus *userbus.Business
	Auth    *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	bearer := mid.Bearer(cfg.Auth)
	basic := mid.Basic(cfg.UserBus, cfg.Auth)

	api := newAPI(cfg.Auth)
	app.HandlerFunc(http.MethodGet, version, "/auth/token/{kid}", api.token, basic)
	app.HandlerFunc(http.MethodGet, version, "/auth/authenticate", api.authenticate, bearer)
	app.HandlerFunc(http.MethodPost, version, "/auth/authorize", api.authorize)
}

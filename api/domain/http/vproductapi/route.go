package vproductapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/api/sdk/http/mid"
	"github.com/David-Kalashir/crs-front/app/domain/vproductapp"
	"github.com/David-Kalashir/crs-front/app/sdk/auth"
	"github.com/David-Kalashir/crs-front/app/sdk/authclient"
	"github.com/David-Kalashir/crs-front/business/domain/userbus"
	"github.com/David-Kalashir/crs-front/business/domain/vproductbus"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log         *logger.Logger
	UserBus     *userbus.Business
	VProductBus *vproductbus.Business
	AuthClient  *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI(vproductapp.NewApp(cfg.VProductBus))
	app.HandlerFunc(http.MethodGet, version, "/vproducts", api.query, authen, ruleAdmin)
}

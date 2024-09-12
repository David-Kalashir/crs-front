package tranapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/api/sdk/http/mid"
	"github.com/David-Kalashir/crs-front/app/domain/tranapp"
	"github.com/David-Kalashir/crs-front/app/sdk/auth"
	"github.com/David-Kalashir/crs-front/app/sdk/authclient"
	"github.com/David-Kalashir/crs-front/business/domain/productbus"
	"github.com/David-Kalashir/crs-front/business/domain/userbus"
	"github.com/David-Kalashir/crs-front/business/sdk/sqldb"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	DB         *sqlx.DB
	UserBus    *userbus.Business
	ProductBus *productbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	transaction := mid.BeginCommitRollback(cfg.Log, sqldb.NewBeginner(cfg.DB))
	ruleAdmin := mid.Authorize(cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI(tranapp.NewApp(cfg.UserBus, cfg.ProductBus))
	app.HandlerFunc(http.MethodPost, version, "/tranexample", api.create, authen, ruleAdmin, transaction)
}

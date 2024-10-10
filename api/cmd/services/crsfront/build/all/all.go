// Package all binds all the routes into the specified app.
package all

import (
	"github.com/David-Kalashir/crs-front/api/domain/http/checkapi"
	"github.com/David-Kalashir/crs-front/api/domain/http/minioapi"
	"github.com/David-Kalashir/crs-front/api/domain/http/rawapi"
	"github.com/David-Kalashir/crs-front/api/domain/http/uisearchpageapi"
	"github.com/David-Kalashir/crs-front/api/sdk/http/mux"
	"github.com/David-Kalashir/crs-front/business/domain/authbus"
	"github.com/David-Kalashir/crs-front/business/domain/uiloginbus"
	"github.com/David-Kalashir/crs-front/business/domain/uiloginbus/templs/logintempl"
	"github.com/David-Kalashir/crs-front/business/domain/uinavbarbus"
	"github.com/David-Kalashir/crs-front/business/domain/uinavbarbus/templs/navbartempl"
	"github.com/David-Kalashir/crs-front/business/domain/uinavbarbus/templs/profiletempl"
	"github.com/David-Kalashir/crs-front/business/domain/uisearchpagebus"
	"github.com/David-Kalashir/crs-front/business/domain/uisearchpagebus/templs/searchpagetempl"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	uinavbarbus := uinavbarbus.NewBusiness(navbartempl.NewTemplate(cfg.Log), profiletempl.NewTemplate(cfg.Log), authbus.NewBusiness(cfg.Log))
	uisearchpagebus := uisearchpagebus.NewBusiness(searchpagetempl.NewTemplate(cfg.Log), uinavbarbus, uiloginbus.NewBusiness(logintempl.NewTemplate(cfg.Log)))

	minioapi.Routes(app, minioapi.Config{
		Log:         cfg.Log,
		MinioClient: cfg.Minio,
	})

	uisearchpageapi.Routes(app, uisearchpageapi.Config{
		Log:             cfg.Log,
		UIsearchpagebus: uisearchpagebus,
		StoreKey:        cfg.StoreKey,
	})

	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})
	rawapi.Routes(app)

}

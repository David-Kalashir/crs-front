package rawapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	const version = "v1"

	app.RawHandlerFunc(http.MethodGet, version, "/raw", rawHandler)
}

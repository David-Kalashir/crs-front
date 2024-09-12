// Package checkapi maintains the web based api for system access.
package checkapi

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/checkapp"
	"github.com/David-Kalashir/crs-front/app/sdk/errs"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

type api struct {
	checkApp *checkapp.App
}

func newAPI(checkApp *checkapp.App) *api {
	return &api{
		checkApp: checkApp,
	}
}

// readiness checks if the database is ready and if not will return a 500 status.
// Do not respond by just returning an error because further up in the call
// stack it will interpret that as a non-trusted error.
func (api *api) readiness(ctx context.Context, r *http.Request) web.Encoder {
	if err := api.checkApp.Readiness(ctx); err != nil {
		return errs.NewError(err)
	}

	return ready{Status: "OK"}
}

// liveness returns simple status info if the service is alive. If the
// app is deployed to a Kubernetes cluster, it will also return pod, node, and
// namespace details via the Downward API. The Kubernetes environment variables
// need to be set within your Pod/Deployment manifest.
func (api *api) liveness(ctx context.Context, r *http.Request) web.Encoder {
	return api.checkApp.Liveness()
}

package mid

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/sdk/mid"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Panics executes the panic middleware functionality.
func Panics() web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) mid.Encoder {
		return mid.Panics(ctx, next)
	}

	return addMidFunc(midFunc)
}

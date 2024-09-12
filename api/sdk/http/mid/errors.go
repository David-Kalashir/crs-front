package mid

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/sdk/mid"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Errors executes the errors middleware functionality.
func Errors(log *logger.Logger) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) mid.Encoder {
		return mid.Errors(ctx, log, next)
	}

	return addMidFunc(midFunc)
}

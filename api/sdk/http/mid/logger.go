package mid

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/sdk/mid"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// Logger executes the logger middleware functionality.
func Logger(log *logger.Logger) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) mid.Encoder {
		return mid.Logger(ctx, log, r.URL.Path, r.URL.RawQuery, r.Method, r.RemoteAddr, next)
	}

	return addMidFunc(midFunc)
}

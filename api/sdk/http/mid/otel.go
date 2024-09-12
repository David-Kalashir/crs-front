package mid

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/sdk/mid"
	"github.com/David-Kalashir/crs-front/foundation/web"
	"go.opentelemetry.io/otel/trace"
)

// Otel executes the otel middleware functionality.
func Otel(tracer trace.Tracer) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) mid.Encoder {
		return mid.Otel(ctx, tracer, next)
	}

	return addMidFunc(midFunc)
}

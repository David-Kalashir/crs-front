package mid

import (
	"context"

	"github.com/David-Kalashir/crs-front/foundation/otel"
	"go.opentelemetry.io/otel/trace"
)

// Otel starts the otel tracing and stores the trace id in the context.
func Otel(ctx context.Context, tracer trace.Tracer, next HandlerFunc) Encoder {
	ctx = otel.InjectTracing(ctx, tracer)

	return next(ctx)
}

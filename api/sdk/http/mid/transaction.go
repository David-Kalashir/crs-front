package mid

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/sdk/mid"
	"github.com/David-Kalashir/crs-front/business/sdk/sqldb"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

// BeginCommitRollback executes the transaction middleware functionality.
func BeginCommitRollback(log *logger.Logger, bgn sqldb.Beginner) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) mid.Encoder {
		return mid.BeginCommitRollback(ctx, log, bgn, next)
	}

	return addMidFunc(midFunc)
}

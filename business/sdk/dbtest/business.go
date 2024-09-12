package dbtest

import (
	"time"

	"github.com/David-Kalashir/crs-front/business/domain/homebus"
	"github.com/David-Kalashir/crs-front/business/domain/homebus/stores/homedb"
	"github.com/David-Kalashir/crs-front/business/domain/productbus"
	"github.com/David-Kalashir/crs-front/business/domain/productbus/stores/productdb"
	"github.com/David-Kalashir/crs-front/business/domain/userbus"
	"github.com/David-Kalashir/crs-front/business/domain/userbus/stores/usercache"
	"github.com/David-Kalashir/crs-front/business/domain/userbus/stores/userdb"
	"github.com/David-Kalashir/crs-front/business/domain/vproductbus"
	"github.com/David-Kalashir/crs-front/business/domain/vproductbus/stores/vproductdb"
	"github.com/David-Kalashir/crs-front/business/sdk/delegate"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// BusDomain represents all the business domain apis needed for testing.
type BusDomain struct {
	Delegate *delegate.Delegate
	Home     *homebus.Business
	Product  *productbus.Business
	User     *userbus.Business
	VProduct *vproductbus.Business
}

func newBusDomains(log *logger.Logger, db *sqlx.DB) BusDomain {
	delegate := delegate.New(log)
	userBus := userbus.NewBusiness(log, delegate, usercache.NewStore(log, userdb.NewStore(log, db), time.Hour))
	productBus := productbus.NewBusiness(log, userBus, delegate, productdb.NewStore(log, db))
	homeBus := homebus.NewBusiness(log, userBus, delegate, homedb.NewStore(log, db))
	vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(log, db))

	return BusDomain{
		Delegate: delegate,
		Home:     homeBus,
		Product:  productBus,
		User:     userBus,
		VProduct: vproductBus,
	}
}

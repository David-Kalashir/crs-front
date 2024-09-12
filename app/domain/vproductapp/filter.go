package vproductapp

import (
	"strconv"

	"github.com/David-Kalashir/crs-front/app/sdk/errs"
	"github.com/David-Kalashir/crs-front/business/domain/vproductbus"
	"github.com/David-Kalashir/crs-front/business/types/name"
	"github.com/google/uuid"
)

func parseFilter(qp QueryParams) (vproductbus.QueryFilter, error) {
	var filter vproductbus.QueryFilter

	if qp.ID != "" {
		id, err := uuid.Parse(qp.ID)
		if err != nil {
			return vproductbus.QueryFilter{}, errs.NewFieldsError("product_id", err)
		}
		filter.ID = &id
	}

	if qp.Name != "" {
		name, err := name.Parse(qp.Name)
		if err != nil {
			return vproductbus.QueryFilter{}, errs.NewFieldsError("name", err)
		}
		filter.Name = &name
	}

	if qp.Cost != "" {
		cst, err := strconv.ParseFloat(qp.Cost, 64)
		if err != nil {
			return vproductbus.QueryFilter{}, errs.NewFieldsError("cost", err)
		}
		filter.Cost = &cst
	}

	if qp.Quantity != "" {
		qua, err := strconv.ParseInt(qp.Quantity, 10, 64)
		if err != nil {
			return vproductbus.QueryFilter{}, errs.NewFieldsError("quantity", err)
		}
		i := int(qua)
		filter.Quantity = &i
	}

	if qp.Name != "" {
		name, err := name.Parse(qp.Name)
		if err != nil {
			return vproductbus.QueryFilter{}, errs.NewFieldsError("name", err)
		}
		filter.UserName = &name
	}

	return filter, nil
}

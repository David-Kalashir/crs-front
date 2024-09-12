package vproductapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/vproductapp"
)

func parseQueryParams(r *http.Request) vproductapp.QueryParams {
	values := r.URL.Query()

	filter := vproductapp.QueryParams{
		Page:     values.Get("page"),
		Rows:     values.Get("row"),
		OrderBy:  values.Get("orderBy"),
		ID:       values.Get("product_id"),
		Name:     values.Get("name"),
		Cost:     values.Get("cost"),
		Quantity: values.Get("quantity"),
		UserName: values.Get("user_name"),
	}

	return filter
}

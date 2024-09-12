package productapi

import (
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/productapp"
)

func parseQueryParams(r *http.Request) productapp.QueryParams {
	values := r.URL.Query()

	filter := productapp.QueryParams{
		Page:     values.Get("page"),
		Rows:     values.Get("row"),
		OrderBy:  values.Get("orderBy"),
		ID:       values.Get("product_id"),
		Name:     values.Get("name"),
		Cost:     values.Get("cost"),
		Quantity: values.Get("quantity"),
	}

	return filter
}

// Package vproductapi maintains the web based api for product view access.
package vproductapi

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/vproductapp"
	"github.com/David-Kalashir/crs-front/app/sdk/errs"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

type api struct {
	vproductApp *vproductapp.App
}

func newAPI(vproductApp *vproductapp.App) *api {
	return &api{
		vproductApp: vproductApp,
	}
}

func (api *api) query(ctx context.Context, r *http.Request) web.Encoder {
	qp := parseQueryParams(r)

	prd, err := api.vproductApp.Query(ctx, qp)
	if err != nil {
		return errs.NewError(err)
	}

	return prd
}

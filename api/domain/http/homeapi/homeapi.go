// Package homeapi maintains the web based api for home access.
package homeapi

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/homeapp"
	"github.com/David-Kalashir/crs-front/app/sdk/errs"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

type api struct {
	homeApp *homeapp.App
}

func newAPI(homeApp *homeapp.App) *api {
	return &api{
		homeApp: homeApp,
	}
}

func (api *api) create(ctx context.Context, r *http.Request) web.Encoder {
	var app homeapp.NewHome
	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	hme, err := api.homeApp.Create(ctx, app)
	if err != nil {
		return errs.NewError(err)
	}

	return hme
}

func (api *api) update(ctx context.Context, r *http.Request) web.Encoder {
	var app homeapp.UpdateHome
	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	hme, err := api.homeApp.Update(ctx, app)
	if err != nil {
		return errs.NewError(err)
	}

	return hme
}

func (api *api) delete(ctx context.Context, r *http.Request) web.Encoder {
	if err := api.homeApp.Delete(ctx); err != nil {
		return errs.NewError(err)
	}

	return nil
}

func (api *api) query(ctx context.Context, r *http.Request) web.Encoder {
	qp := parseQueryParams(r)

	hme, err := api.homeApp.Query(ctx, qp)
	if err != nil {
		return errs.NewError(err)
	}

	return hme
}

func (api *api) queryByID(ctx context.Context, r *http.Request) web.Encoder {
	hme, err := api.homeApp.QueryByID(ctx)
	if err != nil {
		return errs.NewError(err)
	}

	return hme
}

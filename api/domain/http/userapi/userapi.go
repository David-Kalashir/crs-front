// Package userapi maintains the web based api for user access.
package userapi

import (
	"context"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/userapp"
	"github.com/David-Kalashir/crs-front/app/sdk/errs"
	"github.com/David-Kalashir/crs-front/foundation/web"
)

type api struct {
	userApp *userapp.App
}

func newAPI(userApp *userapp.App) *api {
	return &api{
		userApp: userApp,
	}
}

func (api *api) create(ctx context.Context, r *http.Request) web.Encoder {
	var app userapp.NewUser
	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	usr, err := api.userApp.Create(ctx, app)
	if err != nil {
		return errs.NewError(err)
	}

	return usr
}

func (api *api) update(ctx context.Context, r *http.Request) web.Encoder {
	var app userapp.UpdateUser
	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	usr, err := api.userApp.Update(ctx, app)
	if err != nil {
		return errs.NewError(err)
	}

	return usr
}

func (api *api) updateRole(ctx context.Context, r *http.Request) web.Encoder {
	var app userapp.UpdateUserRole
	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	usr, err := api.userApp.UpdateRole(ctx, app)
	if err != nil {
		return errs.NewError(err)
	}

	return usr
}

func (api *api) delete(ctx context.Context, r *http.Request) web.Encoder {
	if err := api.userApp.Delete(ctx); err != nil {
		return errs.NewError(err)
	}

	return nil
}

func (api *api) query(ctx context.Context, r *http.Request) web.Encoder {
	qp, err := parseQueryParams(r)
	if err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	usr, err := api.userApp.Query(ctx, qp)
	if err != nil {
		return errs.NewError(err)
	}

	return usr
}

func (api *api) queryByID(ctx context.Context, r *http.Request) web.Encoder {
	usr, err := api.userApp.QueryByID(ctx)
	if err != nil {
		return errs.NewError(err)
	}

	return usr
}

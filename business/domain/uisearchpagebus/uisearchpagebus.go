package uisearchpagebus

import (
	"context"

	"github.com/David-Kalashir/crs-front/business/domain/uiloginbus"
	"github.com/David-Kalashir/crs-front/business/domain/uinavbarbus"
	"github.com/a-h/templ"
	"github.com/google/uuid"
)

// Storer interface declares the behavior this package needs to persist and
// retrieve data.
type Templer interface {
	SearchPage(ctx context.Context, navbar Component) templ.Component
}

// Business manages the set of APIs for view product access.
type Business struct {
	templer     Templer
	uinavbarbus *uinavbarbus.Business
	uiloginbus  *uiloginbus.Business
}

// NewBusiness constructs a vproduct business API for use.
func NewBusiness(templer Templer, uinavbarbus *uinavbarbus.Business, uiloginbus *uiloginbus.Business) *Business {
	return &Business{
		templer:     templer,
		uinavbarbus: uinavbarbus,
		uiloginbus:  uiloginbus,
	}
}

// SearchPage .
func (b *Business) SearchPage(ctx context.Context, uuid uuid.UUID) (templ.Component, error) {
	navabr, err := b.uinavbarbus.Navbar(ctx, "/v1/searchpage/login", "/v1/searchpage/logout", uuid, "#login", "body")
	if err != nil {
		return nil, err
	}
	com := b.templer.SearchPage(ctx, Component{
		Component: navabr,
		ID:        "navbar-serach",
	})

	return com, nil
}

func (b *Business) Login(ctx context.Context) (templ.Component, error) {
	login, err := b.uiloginbus.Login(ctx)
	if err != nil {
		return nil, err
	}

	return login, nil
}

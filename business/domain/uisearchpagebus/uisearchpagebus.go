package uisearchpagebus

import (
	"context"

	"github.com/David-Kalashir/crs-front/business/domain/uinavbarbus"
	"github.com/a-h/templ"
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
}

// NewBusiness constructs a vproduct business API for use.
func NewBusiness(templer Templer, uinavbarbus *uinavbarbus.Business) *Business {
	return &Business{
		templer:     templer,
		uinavbarbus: uinavbarbus,
	}
}

// SearchPage .
func (b *Business) SearchPage(ctx context.Context) (templ.Component, error) {
	navabr, err := b.uinavbarbus.Navbar(ctx)
	if err != nil {
		return nil, err
	}
	com := b.templer.SearchPage(ctx, Component{
		Component: navabr,
		ID:        "navbar-serach",
	})

	return com, nil
}

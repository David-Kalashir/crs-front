package uisearchpagebus

import (
	"context"

	"github.com/a-h/templ"
)

// Storer interface declares the behavior this package needs to persist and
// retrieve data.
type Templer interface {
	SearchPage(ctx context.Context) templ.Component
}

// Business manages the set of APIs for view product access.
type Business struct {
	templer Templer
}

// NewBusiness constructs a vproduct business API for use.
func NewBusiness(templer Templer) *Business {
	return &Business{
		templer: templer,
	}
}

// SearchPage .
func (b *Business) SearchPage(ctx context.Context) (templ.Component, error) {
	com := b.templer.SearchPage(ctx)

	return com, nil
}

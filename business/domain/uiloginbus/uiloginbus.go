package uiloginbus

import (
	"context"

	"github.com/a-h/templ"
)

type Templer interface {
	Login(ctx context.Context) templ.Component
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

// Login .
func (b *Business) Login(ctx context.Context) (templ.Component, error) {
	com := b.templer.Login(ctx)

	return com, nil
}

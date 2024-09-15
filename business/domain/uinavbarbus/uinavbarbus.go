package uinavbarbus

import (
	"context"

	"github.com/a-h/templ"
)

// Storer interface declares the behavior this package needs to persist and
// retrieve data.
type NavbarTempler interface {
	Navbar(ctx context.Context, navbarstyl Navbarstyl, c NavbarComponent, id string) templ.Component
}

type ProfileTempler interface {
	Profile(ctx context.Context, id string) templ.Component
}

// Business manages the set of APIs for view product access.
type Business struct {
	navbartempler  NavbarTempler
	profiletempler ProfileTempler
}

// NewBusiness constructs a vproduct business API for use.
func NewBusiness(navbartempler NavbarTempler, profiletempler ProfileTempler) *Business {
	return &Business{
		navbartempler:  navbartempler,
		profiletempler: profiletempler,
	}
}

// SearchPage .
func (b *Business) Navbar(ctx context.Context) (templ.Component, error) {
	p := b.profiletempler.Profile(ctx, "profile")
	com := b.navbartempler.Navbar(ctx, Navbarstyl{
		Background: "",
		Logo:       "",
	}, NavbarComponent{
		Profile: Component{
			Component: p,
			ID:        "profile-navbar",
		},
	}, "")

	return com, nil
}

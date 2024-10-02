package uinavbarbus

import (
	"context"

	"github.com/David-Kalashir/crs-front/business/domain/authbus"
	"github.com/a-h/templ"
	"github.com/google/uuid"
)

// Storer interface declares the behavior this package needs to persist and
// retrieve data.
type NavbarTempler interface {
	Navbar(ctx context.Context, navbarstyl Navbarstyl, c NavbarComponent, id string, route string, authbtn string, authtarget string) templ.Component
}

type ProfileTempler interface {
	Profile(ctx context.Context, id string) templ.Component
}

// Business manages the set of APIs for view product access.
type Business struct {
	navbartempler  NavbarTempler
	profiletempler ProfileTempler
	authbus        *authbus.Business
}

// NewBusiness constructs a vproduct business API for use.
func NewBusiness(navbartempler NavbarTempler, profiletempler ProfileTempler, authbus *authbus.Business) *Business {
	return &Business{
		navbartempler:  navbartempler,
		profiletempler: profiletempler,
		authbus:        authbus,
	}
}

// SearchPage .
func (b *Business) Navbar(ctx context.Context, loginroute, logoutroute string, uuid uuid.UUID, logintarget string, logouttarget string) (templ.Component, error) {
	p := b.profiletempler.Profile(ctx, "profile")
	auth := b.authbus.IsAuth(uuid)

	var com templ.Component

	if auth {
		com = b.navbartempler.Navbar(ctx, Navbarstyl{
			Background: "",
			Logo:       "",
		}, NavbarComponent{
			Profile: Component{
				Component: p,
				ID:        "profile-navbar",
			},
		}, "", logoutroute, "logout", logouttarget)
	}
	if !auth {
		com = b.navbartempler.Navbar(ctx, Navbarstyl{
			Background: "",
			Logo:       "",
		}, NavbarComponent{
			Profile: Component{
				Component: p,
				ID:        "profile-navbar",
			},
		}, "", loginroute, "login", logintarget)
	}

	return com, nil
}

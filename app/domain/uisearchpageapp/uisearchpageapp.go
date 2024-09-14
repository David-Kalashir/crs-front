package uisearchpageapp

import (
	"context"

	"github.com/David-Kalashir/crs-front/business/domain/uisearchpagebus"
	"github.com/a-h/templ"
)

// App manages the set of app layer api functions for the user domain.
type App struct {
	uisearchpagebus *uisearchpagebus.Business
}

// NewApp constructs a user app API for use.
func NewApp(uisearchpagebus *uisearchpagebus.Business) *App {
	return &App{
		uisearchpagebus: uisearchpagebus,
	}
}

// Create adds a new user to the system.
func (a *App) SearchPage(ctx context.Context) (templ.Component, error) {
	com, err := a.uisearchpagebus.SearchPage(ctx)
	if err != nil {
		return nil, err
	}
	return com, nil
}

package searchpagetempl

import (
	"context"

	"github.com/David-Kalashir/crs-front/business/domain/uisearchpagebus"
	"github.com/David-Kalashir/crs-front/foundation/logger"
	"github.com/a-h/templ"
)

// template manages the set of APIs for wallet and transaction database access.
type Template struct {
	log *logger.Logger
}

// NewTemplate constructs the api for data access.
func NewTemplate(log *logger.Logger) *Template {
	return &Template{
		log: log,
	}
}

func (t *Template) SearchPage(ctx context.Context, navbar uisearchpagebus.Component) templ.Component {
	s := SearchTempl{
		Navbar: Component{
			Component: navbar.Component,
			ID:        navbar.ID,
		},
	}
	templ := s.Searchpage()
	return templ
}

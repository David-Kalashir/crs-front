package navbartempl

import (
	"context"

	"github.com/David-Kalashir/crs-front/business/domain/uinavbarbus"
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

func (t *Template) Navbar(ctx context.Context, navbarstyl uinavbarbus.Navbarstyl, c uinavbarbus.NavbarComponent, id string, route string, authbtn string, authtarget string) templ.Component {
	s := NavbarTempl{
		Profile: Component{
			Componenet: c.Profile.Component,
			ID:         c.Profile.ID,
		},
	}
	templ := s.Navbar(Navbarstyl{
		Bg:   navbarstyl.Background,
		Logo: navbarstyl.Background,
	}, id, route, authbtn, authtarget)
	return templ
}

package logintempl

import (
	"context"

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

func (t *Template) Login(ctx context.Context) templ.Component {

	l := LoginTemplate{}

	templ := l.Login()
	return templ
}

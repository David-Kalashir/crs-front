package searchpagetempl

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

func (t *Template) SearchPage(ctx context.Context) templ.Component {
	s := SearchTempl{}
	templ := s.Searchpage()
	return templ
}

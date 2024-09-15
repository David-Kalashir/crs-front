package searchpagetempl

import "github.com/a-h/templ"

type SearchTempl struct {
	Navbar Component
}
type Component struct {
	Component templ.Component
	ID        string
}

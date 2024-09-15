package navbartempl

import "github.com/a-h/templ"

type Navbarstyl struct {
	Bg   string
	Logo string
}

type NavbarTempl struct {
	Profile Component
}
type Component struct {
	Componenet templ.Component
	ID         string
}

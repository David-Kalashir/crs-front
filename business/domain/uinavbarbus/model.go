package uinavbarbus

import "github.com/a-h/templ"

type Navbarstyl struct {
	Background string
	Logo       string
}

type Component struct {
	Component templ.Component
	ID        string
}
type NavbarComponent struct {
	Profile Component
}

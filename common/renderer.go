package common

// Renderable defines objects that can render themselves
type Renderable interface {
	// Render returns object as string
	Render() string
}

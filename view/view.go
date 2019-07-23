package view

import (
	"termui/v3/widgets"

	ui "github.com/gizak/termui/v3"
)

var listeners []string

// View is a set of elements of game UI
type View struct {
	Location  *widgets.Paragraph
	CombatLog *widgets.Paragraph
}

// NewView returns new predefined views
func NewView() *View {
	var view = View{}

	view.Location = widgets.NewParagraph()
	view.Location.Title = "Location"
	view.Location.SetRect(10, 0, 32, 12)

	view.CombatLog = widgets.NewParagraph()
	view.CombatLog.Title = "Combat log"
	view.CombatLog.SetRect(0, 13, 50, 18)

	return &view
}

// UpdateLocation view text
func (v *View) UpdateLocation(text string) {
	v.Location.Text = text
}

// UpdateCombatLog view text
func (v *View) UpdateCombatLog(text string) {
	v.CombatLog.Text = text
}

// All return all available view parts
func (v *View) All() []ui.Drawable {
	return []ui.Drawable{
		v.Location,
		v.CombatLog,
	}
}

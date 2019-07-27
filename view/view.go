package view

import (
	"github.com/gizak/termui/v3/widgets"

	ui "github.com/gizak/termui/v3"
)

// View is a set of elements of game UI
type View struct {
	Location  *widgets.Paragraph
	CombatLog *widgets.Paragraph
	Hero      *widgets.Table
	Monster   *widgets.Table
}

// NewView returns new predefined views
func NewView() *View {
	var view = View{}

	view.Location = widgets.NewParagraph()
	view.Location.Title = "Location"
	view.Location.SetRect(0, 0, 38, 20)

	view.CombatLog = widgets.NewParagraph()
	view.CombatLog.Title = "Combat log"
	view.CombatLog.SetRect(0, 20, 55, 25)

	view.Hero = widgets.NewTable()
	view.Hero.Title = "My hero"
	view.Hero.SetRect(40, 0, 55, 5)

	view.Monster = widgets.NewTable()
	view.Monster.Title = ""
	view.Monster.Rows = [][]string{[]string{""}}
	view.Monster.SetRect(40, 5, 55, 10)

	return &view
}

// UpdateLocation view change
func (v *View) UpdateLocation(text string) {
	v.Location.Text = text
}

// UpdateCombatLog view change
func (v *View) UpdateCombatLog(text string) {
	v.CombatLog.Text = text
}

// UpdateHeroStats view change
func (v *View) UpdateHeroStats(stats [][]string) {
	v.Hero.Rows = stats
}

// ShowMonster creates new view for current monster
func (v *View) ShowMonster(stats [][]string) {
	v.Monster.Title = "Monster"
	v.Monster.Rows = stats
}

// HideMonster removes monster from view
func (v *View) HideMonster() {
	v.Monster.Title = ""
	v.Monster.Rows = [][]string{[]string{""}}
}

// All return all available view parts
func (v *View) All() []ui.Drawable {
	return []ui.Drawable{
		v.Location,
		v.CombatLog,
		v.Hero,
		v.Monster,
	}
}

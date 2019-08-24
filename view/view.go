package view

import (
	"fmt"
	"github.com/gizak/termui/v3/widgets"

	ui "github.com/gizak/termui/v3"
)

// View is a set of elements of game UI
type View struct {
	Goal      *widgets.Paragraph
	Location  *widgets.Paragraph
	CombatLog *widgets.Paragraph
	Hero      *widgets.Table
	SkillsBar *widgets.Table
	Monster   *widgets.Table
}

// NewView returns new predefined views
func NewView() *View {
	var view = View{}

	view.Goal = widgets.NewParagraph()
	view.Goal.Title = "Your goal"
	view.Goal.SetRect(0, 0, 65, 3)

	view.Location = widgets.NewParagraph()
	view.Location.Title = "Location"
	view.Location.SetRect(0, 3, 38, 23)

	view.SkillsBar = widgets.NewTable()
	view.SkillsBar.Title = "Skill bar"
	view.SkillsBar.Rows = [][]string{{""}}
	view.SkillsBar.SetRect(0, 23, 65, 28)

	view.CombatLog = widgets.NewParagraph()
	view.CombatLog.Title = "Combat log"
	view.CombatLog.SetRect(0, 28, 65, 34)

	view.Hero = widgets.NewTable()
	view.Hero.Title = "My hero"
	view.Hero.Rows = [][]string{{""}}
	view.Hero.SetRect(40, 3, 65, 14)

	view.Monster = widgets.NewTable()
	view.Monster.Title = ""
	view.Monster.Rows = [][]string{{""}}
	view.Monster.SetRect(40, 14, 65, 23)

	return &view
}

// UpdateLocation view change
func (v *View) UpdateLocation(text string) {
	v.Location.Text = text
}

// UpdateLocationTitle view change
func (v *View) UpdateLocationTitle(level int) {
	v.Location.Title = "Location level " + fmt.Sprintf("%d", level)
}

func (v *View) UpdateGoal(monster string, kills, time int) {
	v.Goal.Text = fmt.Sprintf("Kill %d [%s] monsters in %d seconds", kills, monster, time)
}

// UpdateCombatLog view change
func (v *View) UpdateCombatLog(text string) {
	v.CombatLog.Text = text
}

// UpdateHeroStats view change
func (v *View) UpdateHeroStats(stats [][]string) {
	v.Hero.Rows = stats
}

// UpdateSkillBar view change
func (v *View) UpdateSkillBar(skills [][]string) {
	v.SkillsBar.Rows = skills
}

// ShowMonster creates new view for current monster
func (v *View) ShowMonster(stats [][]string) {
	v.Monster.Title = "Monster"
	v.Monster.Rows = stats
}

// HideMonster removes monster from view
func (v *View) HideMonster() {
	v.Monster.Title = ""
	v.Monster.Rows = [][]string{{""}}
}

// All return all available view parts
func (v *View) All() []ui.Drawable {
	return []ui.Drawable{
		v.Goal,
		v.Location,
		v.CombatLog,
		v.Hero,
		v.SkillsBar,
		v.Monster,
	}
}

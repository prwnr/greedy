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

// UIChange channel triggered on Views changes
var UIChange = make(chan bool)

// NewView returns new predefined views
func NewView() *View {
	var v = View{}

	v.Goal = widgets.NewParagraph()
	v.Goal.Title = "Your goal"
	v.Goal.SetRect(0, 0, 65, 3)

	v.Location = widgets.NewParagraph()
	v.Location.Title = "Location"
	v.Location.SetRect(0, 3, 38, 23)

	v.SkillsBar = widgets.NewTable()
	v.SkillsBar.Title = "Skill bar"
	v.SkillsBar.Rows = [][]string{{""}}
	v.SkillsBar.SetRect(0, 23, 65, 30)

	v.CombatLog = widgets.NewParagraph()
	v.CombatLog.Title = "Combat log"
	v.CombatLog.SetRect(0, 30, 65, 37)

	v.Hero = widgets.NewTable()
	v.Hero.Title = "My hero"
	v.Hero.Rows = [][]string{{""}}
	v.Hero.SetRect(40, 3, 65, 14)

	v.Monster = widgets.NewTable()
	v.Monster.Title = ""
	v.Monster.Rows = [][]string{{""}}
	v.Monster.SetRect(40, 14, 65, 23)

	ui.Render(v.All()...)

	return &v
}

// UpdateLocation view change
func (v *View) UpdateLocation(text string) {
	v.Location.Text = text
	renderView(v.Location)
}

// UpdateLocationTitle view change
func (v *View) UpdateLocationTitle(level int) {
	v.Location.Title = "Location level " + fmt.Sprintf("%d", level)
	renderView(v.Location)
}

func (v *View) UpdateGoal(monster string, kills, time int) {
	v.Goal.Text = fmt.Sprintf("Kill %d [%s] monsters in %d seconds", kills, monster, time)
	renderView(v.Goal)
}

// UpdateCombatLog view change
func (v *View) UpdateCombatLog(text string) {
	v.CombatLog.Text = text
	renderView(v.CombatLog)
}

// UpdateHeroStats view change
func (v *View) UpdateHeroStats(stats [][]string) {
	v.Hero.Rows = stats
	renderView(v.Hero)
}

// UpdateSkillBar view change
func (v *View) UpdateSkillBar(skills [][]string) {
	v.SkillsBar.Rows = skills
	renderView(v.SkillsBar)
}

// ShowMonster creates new view for current monster
func (v *View) ShowMonster(stats [][]string) {
	v.Monster.Title = "Monster"
	v.Monster.Rows = stats
	renderView(v.Monster)
}

// HideMonster removes monster from view
func (v *View) HideMonster() {
	v.Monster.Title = ""
	v.Monster.Rows = [][]string{{""}}
	renderView(v.Monster)
}

func renderView(v ui.Drawable) {
	ui.Render(v)
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

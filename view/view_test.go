package view

import (
	"reflect"
	"testing"
)

func TestViewsUpdate(t *testing.T) {
	t.Run("test combat log update", func(t *testing.T) {
		v := NewView()

		v.UpdateCombatLog("new log")
		assertStringEquals(t, "new log", v.CombatLog.Text)
	})

	t.Run("test location update", func(t *testing.T) {
		v := NewView()

		v.UpdateLocation("new location")
		assertStringEquals(t, "new location", v.Location.Text)
	})

	t.Run("test hero stats update", func(t *testing.T) {
		v := NewView()

		want := [][]string{
			[]string{"Foo"},
		}
		v.UpdateHeroStats(want)

		got := v.Hero.Rows
		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted stats on vie %v, but got %v", want, got)
		}
	})

	t.Run("test monster update", func(t *testing.T) {
		v := NewView()

		want := [][]string{
			[]string{"Foo"},
		}
		v.ShowMonster(want)

		got := v.Monster.Rows

		assertStringEquals(t, "Monster", v.Monster.Title)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted stats on vie %v, but got %v", want, got)
		}
	})
}

func assertStringEquals(t *testing.T, want, got string) {
	if want != got {
		t.Errorf("wanted string %s, got %s", want, got)
	}
}

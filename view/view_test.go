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

		assertSlicesEqual(t, want, v.Hero.Rows)
	})

	t.Run("test monster update", func(t *testing.T) {
		v := NewView()

		want := [][]string{
			[]string{"Foo"},
		}
		v.ShowMonster(want)


		assertStringEquals(t, "Monster", v.Monster.Title)
		assertSlicesEqual(t, want, v.Monster.Rows)
	})

	t.Run("test hides monster update", func(t *testing.T) {
		v := NewView()

		want := [][]string{
			[]string{"Foo"},
		}
		v.ShowMonster(want)

		assertStringEquals(t, "Monster", v.Monster.Title)
		assertSlicesEqual(t, want, v.Monster.Rows)

		v.HideMonster()
		want = [][]string{[]string{""}}
		assertStringEquals(t, "", v.Monster.Title)
		assertSlicesEqual(t, want, v.Monster.Rows)
	})
}

func assertStringEquals(t *testing.T, want, got string) {
	if want != got {
		t.Errorf("wanted string %s, got %s", want, got)
	}
}

func assertSlicesEqual(t *testing.T, want, got [][]string) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted stats on vie %v, but got %v", want, got)
	}
}

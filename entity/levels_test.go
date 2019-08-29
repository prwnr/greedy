package entity

import (
	"testing"
)

func assertLevelEquals(t *testing.T, want, got int) {
	if want != got {
		t.Errorf("level should be at %d, got %d", want, got)
	}
}

func assertNil(t *testing.T, actual *HeroLevel) {
	if actual != nil {
		t.Errorf("actual value should be nil, got %v", actual)
	}
}

func assertNotNil(t *testing.T, actual *HeroLevel) {
	if actual == nil {
		t.Error("actual value shouldn't be nil")
	}
}

func TestNewHeroLevel(t *testing.T) {
	t.Run("creates single level", func(t *testing.T) {
		l := NewHeroLevel(3, 0)

		got := l.Number
		assertLevelEquals(t, 3, got)
		assertNil(t, l.Next)
	})

	t.Run("creates levels hierarchy", func(t *testing.T) {
		l := NewHeroLevel(1, 5)

		//assert three non-empty levels
		for i := 1; i < 5; i++ {
			assertLevelEquals(t, i, l.Number)

			l = l.Next
			assertNotNil(t, l)
		}

		//assert last level is final
		assertNil(t, l.Next)
	})
}

func TestNewMonsterLevel(t *testing.T) {
	l := NewMonsterLevel(3)
	assertLevelEquals(t, 3, l.Number)

	l = NewMonsterLevel(2)
	assertLevelEquals(t, 2, l.Number)
}

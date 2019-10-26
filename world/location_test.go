package world

import (
	"greedy/entity"
	"testing"
)

func TestLocationCreation(t *testing.T) {
	l := NewLocation(5, 1) // created 2 two monsters

	assertHasFreeSpots(t, l)
	assertCount(t, len(l.Places), 5)
	assertCount(t, len(l.Places[0]), 5)
}

func TestLocationRender(t *testing.T) {
	t.Run("render empty location", func(t *testing.T) {
		l := NewLocation(1, 1)

		got := l.RenderPlaces()
		want := "_\r\n"
		if got != want {
			t.Errorf("rendered location is '%s', wanted '%s'", got, want)
		}
	})

	t.Run("render location with hero", func(t *testing.T) {
		l := NewLocation(1, 1)
		h := entity.NewHero(0, 0)
		l.PlaceHero(h)

		got := l.RenderPlaces()
		want := "*\r\n"
		if got != want {
			t.Errorf("rendered location is '%s', wanted '%s'", got, want)
		}
	})
}

func TestPlacingMonsters(t *testing.T) {
	t.Run("placing few monsters on location", func(t *testing.T) {
		l := NewLocation(5, 1) // created with 2 monsters

		assertHasFreeSpots(t, l)
		assertCount(t, len(l.Places), 5)
		assertCount(t, len(l.Places[0]), 5)

		l.PlaceMonsters(5)
		assertHasFreeSpots(t, l)
	})

	t.Run("placing maximum number of monsters", func(t *testing.T) {
		l := NewLocation(2, 1) // created with 1 monster, 3 spots left

		assertHasFreeSpots(t, l)
		assertCount(t, len(l.Places), 2)
		assertCount(t, len(l.Places[0]), 2)

		l.PlaceMonsters(3)
		assertNotHasFreeSpots(t, l)
	})

	t.Run("placing more monster that location can contain wont work", func(t *testing.T) {
		l := NewLocation(2, 1) // created with 1 monster, 3 spots left

		assertHasFreeSpots(t, l)
		assertCount(t, len(l.Places), 2)
		assertCount(t, len(l.Places[0]), 2)

		l.PlaceMonsters(5)
		assertNotHasFreeSpots(t, l)
	})

	t.Run("placing new monster after removing one", func(t *testing.T) {
		l := NewLocation(2, 1) // created with 1 monster, 3 spots left

		assertHasFreeSpots(t, l)
		assertCount(t, len(l.Places), 2)
		assertCount(t, len(l.Places[0]), 2)

		l.PlaceMonsters(3)
		assertNotHasFreeSpots(t, l)

		for i := 0; i < 2; i++ {
			l.Places[i][i].RemoveMonster()
			assertHasFreeSpots(t, l)

			l.PlaceMonsters(1)
			assertNotHasFreeSpots(t, l)
		}
	})

	t.Run("location monsters are scaling with location level", func(t *testing.T) {
		assertBetween := func(t *testing.T, min, max, got int) {
			if got < min || got > max {
				t.Errorf("want level between %d and %d, got %d", min, max, got)
			}
		}

		var l *Location

		l = NewLocation(1, 1)
		l.PlaceMonsters(1)
		monster := l.Places[0][0].GetMonster()
		assertBetween(t, 1, 3, monster.Level())

		l = NewLocation(1, 3)
		l.PlaceMonsters(1)
		assertBetween(t, 3, 5, l.Places[0][0].GetMonster().Level())
	})
}

func TestLocationWithHero(t *testing.T) {
	assertHeroIsOnPosition := func(t *testing.T, l *Location, x, y int) {
		if l.Places[y][x].GetHero() == nil {
			t.Errorf("hero should be on position X%d, Y%d, but is not found", x, y)
		}
	}

	t.Run("placing a hero on location", func(t *testing.T) {
		l := NewLocation(2, 1) // created with 1 monster, 3 spots left
		h := entity.NewHero(0, 1)

		l.PlaceHero(h)

		assertHasFreeSpots(t, l)
		assertHeroIsOnPosition(t, l, 0, 1)
	})

	t.Run("changing hero place", func(t *testing.T) {
		l := NewLocation(2, 1) // created with 1 monster, 3 spots left
		h := entity.NewHero(0, 1)

		l.PlaceHero(h)

		assertHasFreeSpots(t, l)
		assertHeroIsOnPosition(t, l, 0, 1)

		h.Position.X = 1

		l.PlaceHero(h)
		assertHeroIsOnPosition(t, l, 1, 1)
	})
}

func assertHasFreeSpots(t *testing.T, l *Location) {
	if !l.HasFreePlace() {
		t.Errorf("location should have free spots")
	}
}

func assertNotHasFreeSpots(t *testing.T, l *Location) {
	if l.HasFreePlace() {
		t.Errorf("location shouldn't have free spots")
	}
}

func assertCount(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d elements, want %d", got, want)
	}
}

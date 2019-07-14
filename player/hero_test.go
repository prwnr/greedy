package player

import "testing"

func TestSetsStartingHeroPosition(t *testing.T) {
	h := Hero{}
	h.Start(1, 2)

	if h.Position.X != 1 || h.Position.Y != 2 {
		t.Errorf("got wrong hero position, expected 1;2. got %d;%d", h.Position.X, h.Position.Y)
	}
}

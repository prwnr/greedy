package world

import (
	"strings"
	"swarm/player"
	"testing"
)

func TestPlaceAvatars(t *testing.T) {
	t.Run("can set new Avatar on place", func(t *testing.T) {
		p := Place{}
		h := player.Hero{}

		p.AddAvatar(h)

		got := len(p.GetAvatars())
		if got != 1 {
			t.Errorf("expected 1 avatar, got %d", got)
		}
	})

	t.Run("can remove all Avatars from place", func(t *testing.T) {
		p := Place{}
		h := player.Hero{}

		p.AddAvatar(h)
		p.RemoveAvatars()

		got := len(p.GetAvatars())
		if got != 0 {
			t.Errorf("expected 0 avatars, got %d", got)
		}
	})
}

func TestPlaceRendering(t *testing.T) {
	t.Run("can render Avatars from place", func(t *testing.T) {
		p := Place{}
		h := player.Hero{}

		p.AddAvatar(h)

		got := p.Render()
		if strings.Compare(got, "*") != 0 {
			t.Errorf("expected to have '*' Avatar rendered, got %s", got)
		}
	})

	t.Run("renders empy place when there are no Avatars", func(t *testing.T) {
		p := Place{}

		got := p.Render()
		if strings.Compare(got, "_") != 0 {
			t.Errorf("expected to have empty place '_' rendered, got %s", got)
		}
	})
}

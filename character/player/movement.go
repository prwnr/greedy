package player

// MoveUp lowers Y position
func (h *Hero) MoveUp() {
	if h.Position.Y > 0 {
		h.Position.Y--
	}
}

// MoveDown raises Y position
func (h *Hero) MoveDown(max int) {
	if h.Position.Y < max {
		h.Position.Y++
	}
}

// MoveLeft lowers X position
func (h *Hero) MoveLeft() {
	if h.Position.X > 0 {
		h.Position.X--
	}
}

// MoveRight raises X position
func (h *Hero) MoveRight(max int) {
	if h.Position.X < max {
		h.Position.X++
	}
}
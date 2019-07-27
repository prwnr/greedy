package common

import "testing"

func TestGeneratingRandomNumber(t *testing.T) {
	t.Run("generates number no higher than max", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			num := RandomNumber(100)
			if num > 100 {
				t.Errorf("got number higher than 100, where it shouldnt %d", num)
			}
		}
	})

	t.Run("generates different number", func(t *testing.T) {
		first := RandomNumber(100)
		second := RandomNumber(100)

		if first == second {
			t.Error("first and second number were equal and they shouldnt")
		}
	})
}

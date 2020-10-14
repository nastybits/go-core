package fibonacci

import "testing"

func TestNumber(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{32, 1346269},
	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			got, err := Number(tt.arg)
			if (tt.arg <= 0 || tt.arg > 20) && err == nil {
				t.Errorf("Number(%d) = %d, ожидалось сообщение об ошибке", tt.arg, got)
			}
			if tt.arg > 0 && got != tt.want {
				t.Errorf("Number(%d) = %d, ожидалось %d", tt.arg, got, tt.want)
			}
		})
	}
}

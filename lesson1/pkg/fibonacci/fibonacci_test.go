package fibonacci

import "testing"

func TestNum(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{"1", -1, 0},
		{"2", 0, 0},
		{"3", 1, 0},
		{"4", 2, 1},
		{"5", 3, 1},
		{"6", 4, 2},
		{"7", 32, 1346269},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Num(tt.arg)
			if tt.arg > 0 && got != tt.want {
				t.Errorf("Number(%d) = %d, ожидалось %d", tt.arg, got, tt.want)
			}
		})
	}
}

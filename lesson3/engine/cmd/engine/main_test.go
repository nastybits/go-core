package main

import (
	"go.core/lesson3/engine/pkg/stub"
	"testing"
)

func Test_find(t *testing.T) {
	var s stub.Spider
	data, err := scan(s, "stub", 2)
	if err != nil {
		t.Errorf("Ошибка при использовании заглушки")
	}

	tests := []struct {
		name string
		arg  string
		want int
	}{
		{"1", "1", 1},
		{"2", "2", 1},
		{"3", "3", 1},
		{"4", "test", 9},
		{"5", "http", 9},
		{"6", "123", 0},
		{"6", "some", 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := find(tt.arg, data)
			if len(got) != tt.want {
				t.Errorf("Find(%s) = %d, ожидалось %d", tt.arg, len(got), tt.want)
			}
		})
	}
}

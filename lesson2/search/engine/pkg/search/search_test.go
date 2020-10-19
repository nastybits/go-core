package search

import "testing"

func TestFind(t *testing.T) {
	data := make(map[string]string)
	data["http://test.com"] = "Test dot com title"
	data["http://some.com"] = "Some dot com title"
	data["http://test.ru"] 	= "Test dot ru title"
	data["http://some.ru"] 	= "Some dot ru title"
	data["http://test.com/level2"] = "Test dot com level 2 title"

	tests := []struct {
		name string
		arg string
		want int
	}{
		{"1", "Test", 3},
		{"2", "dot", 5},
		{"3", "DOT", 5},
		{"4", "com", 3},
		{"5", "ru", 2},
		{"6", "title", 5},
		{"7", "123", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Find(tt.arg, data)
			if len(got) != tt.want {
				t.Errorf("Find(%s) = %d, ожидалось %d", tt.arg, len(got), tt.want)
			}
		})
	}
}

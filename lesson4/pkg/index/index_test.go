package index

import (
	"go.core/lesson4/pkg/document"
	"testing"
)

func Test_inArray(t *testing.T) {
	type args struct {
		slice []int
		val   int
	}

	tests := []struct {
		name string
		args args
		want bool
	} {
		{"1", args{[]int{1, 2, 3, 4, 5}, 5}, true},
		{"2", args{[]int{1, 2, 3, 4}, 5}, false},
		{"3", args{[]int{0, 2, 3, 4, 5}, 0}, true},
		{"4", args{[]int{0, 2, 3, 4, 5}, -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inArray(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Add(t *testing.T) {
	docs := []document.Documenter{
		document.New(1, "url", "title"),
		document.New(2, "url2", "title2"),
	}

	type args struct {
		d []document.Documenter
	}
	tests := []struct {
		name string
		s    Index
		args args
	}{
		{"1", New(), args{d: docs}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.d[0])
			items, ok := tt.s["url"]
			if !ok || items[0] != 1 {
				t.Errorf("got = %v, want %v", items, 1)
			}
		})
	}
}

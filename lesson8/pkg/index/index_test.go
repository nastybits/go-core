package index

import (
	"go.core/lesson8/pkg/crawler"
	"testing"
)

func TestService_inArray(t *testing.T) {
	s := New()

	type args struct {
		slice []int
		val   int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}, 5}, true},
		{"2", args{[]int{1, 2, 3, 4}, 5}, false},
		{"3", args{[]int{0, 2, 3, 4, 5}, 0}, true},
		{"4", args{[]int{0, 2, 3, 4, 5}, -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.inArray(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Add(t *testing.T) {
	documents := []crawler.Document{
		{1, "url", "title"},
		{2, "url2", "title2"},
	}

	tests := []struct {
		name string
		s    *Index
		docs []crawler.Document
	}{
		{name: "1", s: New(), docs: documents},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.docs[0])
			items, ok := tt.s.data["url"]
			if !ok || items[0] != 1 {
				t.Errorf("got = %v, want %v", items, 1)
			}
		})
	}
}

func TestIndex_Create(t *testing.T) {
	type fields struct {
		data map[string][]int
	}
	type args struct {
		d []crawler.Document
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := &Index{
				data: tt.fields.data,
			}
		})
	}
}

package index

import (
	"go.core/lesson8/pkg/crawler"
	"go.core/lesson8/pkg/fixtures"
	"os"
	"testing"
)

var index *Index

func TestMain(m *testing.M)  {
	index = New()
	os.Exit(m.Run())
}

func TestService_inArray(t *testing.T) {
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
			if got := index.inArray(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Add(t *testing.T) {
	tests := []struct {
		name string
		s    *Index
		docs []crawler.Document
	}{
		{name: "1", s: New(), docs: fixtures.Documents()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index.Add(tt.docs[0])
			items, ok := index.data["url"]
			if !ok || items[0] != 1 {
				t.Errorf("got = %v, want %v", items, 1)
			}
		})
	}
}

func TestIndex_Create(t *testing.T) {
	test := struct {
		name string
		docs []crawler.Document
	}{
		name: "1",
		docs: fixtures.Documents(),
	}

	t.Run(test.name, func(t *testing.T) {
		index.Create(test.docs)
		if (len(index.data)) != 9 {
			t.Errorf("got = %v, want %v", len(index.data), 9)
		}
		items, ok := index.data["two"]
		if !ok || items[0] != 2 {
			t.Errorf("got = %v, want %v", items, 2)
		}
	})
}

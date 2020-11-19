package crawler

import (
	"go.core/lesson6/pkg/crawler/membot"
	"testing"
)


func TestScan(t *testing.T) {
	var s membot.Scanner
	c := New(s)
	tests := []struct {
		name string
		urls []string
		depth int
		want int
	}{
		{name: "One site", urls: []string{"first.site"}, depth: 1, want: 9},
		{name: "Two sites", urls: []string{"first.site", "second.site"}, depth: 2, want: 18},
		{name: "Three sites", urls: []string{"first.site", "second.site", "third.site"}, depth: 3, want: 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			documents, err := c.Scan(tt.urls, tt.depth)
			if err != nil || len(documents) != tt.want {
				t.Errorf("Scan() gotDocuments = %d; want = %d", len(documents), tt.want)
			}
		})
	}
}

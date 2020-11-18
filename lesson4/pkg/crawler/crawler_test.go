package crawler

import (
	"go.core/lesson4/pkg/crawler/stub"
	"testing"
)

var s stub.Scanner

func TestScan(t *testing.T) {
	tests := []struct {
		name string
		urls []string
		depth int
	}{
		{"1", []string{"test.com", "test2.com"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			documents, err := Scan(s, tt.urls, tt.depth)
			if err != nil || documents[0].Title() != "Some test 1 url title" {
				t.Errorf("Scan() gotDocuments = %v", documents)
			}
		})
	}
}

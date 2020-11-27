package crawler

import (
	"go.core/lesson8/pkg/crawler/membot"
	"reflect"
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

func TestDocument_Words(t *testing.T) {
	docs := []Document{
		{ID: 1, URL: "test1.com", Title: "Test 1 one every"},
		{ID: 2, URL: "test2.com", Title: "Test 2 two every in8and2"},
		{ID: 3, URL: "test3.com", Title: "Test 3 three every"},
		{ID: 4, URL: "test4.com", Title: "Test 4 four every in4and5"},
		{ID: 5, URL: "test5.com", Title: "Test 5 five every in4and5"},
		{ID: 6, URL: "test6.com", Title: "Test 6 six every"},
		{ID: 7, URL: "test7.com", Title: "Test 7 seven every"},
		{ID: 8, URL: "test8.com", Title: "Test 8 eight every in8and2"},
		{ID: 9, URL: "test9.com", Title: "Test 9 nine every"},
	}
	tests := []struct {
		name   	string
		doc 	Document
		want   	[]string
	}{
		{name: "Test 1", doc: docs[0], want: []string{ "test1", "com", "test", "1", "one", "every" }},
		{name: "Test 2", doc: docs[1], want: []string{ "test2", "com", "test", "2", "two", "every", "in8and2" }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.doc.Words(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Words() = %v, want %v", got, tt.want)
			}
		})
	}
}

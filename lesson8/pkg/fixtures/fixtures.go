package fixtures

import "go.core/lesson8/pkg/crawler"

func Documents() []crawler.Document {
	docs := []crawler.Document{
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
	return docs
}

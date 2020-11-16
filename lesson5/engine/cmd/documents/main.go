package main

import (
	"fmt"
	"go.core/lesson5/engine/pkg/documents"
)

func main() {
	var d documents.Documents
	docs := []documents.Document{
		{ID: 1, Url: "test1", Title: "Test1"},
		{ID: 2, Url: "test2", Title: "Test2"},
		{ID: 3, Url: "test3", Title: "Test3"},
		{ID: 4, Url: "test4", Title: "Test4"},
		{ID: 5, Url: "test5", Title: "Test5"},
		{ID: 6, Url: "test5", Title: "Test5"},
	}

	for _, doc := range docs {
		d.Insert(doc)
	}

	doc, ok := d.Document(4)
	fmt.Printf("%v\n%v", *doc, ok)
}

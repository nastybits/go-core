package bstree

import (
	"go.core/lesson5/engine/pkg/crawler"
	"reflect"
	"testing"
)

func TestTree_Add(t *testing.T) {
	tree := &Tree{}
	type args struct {
		doc crawler.Document
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Add root element",
			args: args{doc: crawler.Document{ID: 4, URL: "Test4.com", Title: "Test 4"}},
		},
		{
			name: "Add right element",
			args: args{doc: crawler.Document{ID: 5, URL: "Test5.com", Title: "Test 5"}},
		},
		{
			name: "Add left element",
			args: args{doc: crawler.Document{ID: 2, URL: "Test2.com", Title: "Test 2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree.Add(tt.args.doc)
			if tt.args.doc.ID == 4 && tree.root.Doc.ID != 4 {
				t.Errorf("Add() got = %v, want %v", tree.root.Doc.ID, 4)
			}
			if tt.args.doc.ID == 5 && tree.root.right.Doc.ID != 5 {
				t.Errorf("Add() got = %v, want %v", tree.root.right.Doc.ID, 5)
			}
			if tt.args.doc.ID == 2 && tree.root.left.Doc.ID != 2 {
				t.Errorf("Add() got = %v, want %v", tree.root.left.Doc.ID, 2)
			}
		})
	}
}

func TestTree_Document(t *testing.T) {
	tree := &Tree{
		root: &Element{Doc: crawler.Document{ID: 5, URL: "test1.com", Title: "Test 1"}},
	}

	tree.Add(crawler.Document{ID: 5, URL: "test1.com", Title: "Test 5"})
	tree.Add(crawler.Document{ID: 4, URL: "test4.com", Title: "Test 4"})
	tree.Add(crawler.Document{ID: 6, URL: "test6.com", Title: "Test 6"})
	tree.Add(crawler.Document{ID: 9, URL: "test9.com", Title: "Test 9"})

	tests := []struct {
		name  string
		arg  int
		want  crawler.Document
		want1 bool
	}{
		{name: "Search for id 4", arg:  4, want: crawler.Document{ID: 4, URL: "test4.com", Title: "Test 4"}, want1: true},
		{name: "Search for id 10", arg: 10, want: crawler.Document{}, want1: false},
		{name: "Search for id 6", arg: 6, want: crawler.Document{ID: 6, URL: "test6.com", Title: "Test 6"}, want1: true},
		{name: "Search for id 1", arg: 1, want: crawler.Document{}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			got, got1 := tree.Document(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Document() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Document() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

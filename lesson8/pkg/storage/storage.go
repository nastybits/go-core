// Package storage - служба хранения данных
package storage

import (
	"go.core/lesson8/pkg/crawler"
	"go.core/lesson8/pkg/storage/bstree"
)

type Interface interface {
	Create(docs []crawler.Document)
	Document(id int) (crawler.Document, bool)
	Add(d crawler.Document)
}

// New - конструктор службы хранения данных
func New() *bstree.Tree {
	t := bstree.Tree{}
	return &t
}

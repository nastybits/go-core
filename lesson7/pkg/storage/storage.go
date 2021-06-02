// Package storage - служба хранения данных
package storage

import (
	"go.core/lesson7/pkg/crawler"
	"go.core/lesson7/pkg/storage/bstree"
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

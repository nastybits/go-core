// Package storage - служба хранения данных
package storage

import (
	"go.core/lesson6/pkg/crawler"
)

type Storager interface {
	Create(docs []crawler.Document)
	Document(id int) (crawler.Document, bool)
	Add(d crawler.Document)
}

// New - конструктор службы хранения данных
func New(s Storager) Storager {
	return s
}

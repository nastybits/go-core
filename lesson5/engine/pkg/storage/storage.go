// Package storage - служба хранения данных
package storage

import (
	"go.core/lesson5/engine/pkg/crawler"
)

type Storager interface {
	Document(id int) (crawler.Document, bool)
	Add(d crawler.Document)
	Remove(id int) error
}

type Service struct {
	Docs Storager
}

// New - конструктор службы хранения данных
func New(d Storager) Service {
	return Service{
		Docs: d,
	}
}

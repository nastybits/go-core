// Package storage - служба хранения данных
package storage

import (
	"go.core/lesson5/engine/pkg/crawler"
)

type Storager interface {
	Document(id int) (crawler.Document, bool)
	Add(d crawler.Document)
}

type Service struct {
	Storager
}

// New - конструктор службы хранения данных
func New(s Storager) Service {
	return Service{
		s,
	}
}

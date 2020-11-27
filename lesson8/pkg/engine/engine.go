// Package engine - пакет поискового движка
package engine

import (
	"encoding/json"
	"errors"
	"go.core/lesson8/pkg/cache"
	"go.core/lesson8/pkg/crawler"
	"go.core/lesson8/pkg/index"
	"go.core/lesson8/pkg/storage"
)

type Service struct {
	index   index.Interface
	storage storage.Interface
	cache   cache.Interface
}

func New(index index.Interface, storage storage.Interface, cache cache.Interface) *Service {
	s := Service{
		index:   index,
		storage: storage,
		cache:   cache,
	}
	return &s
}

// Load - загружает документы из кэша, индексирует и сохраняет их в хранилище
func (s *Service) Load() error {
	data, err := s.cache.Load()
	if err != nil {
		return errors.New("не удалось загрузить данные из кэша")
	}

	if len(data) == 0 {
		return nil
	}

	var docs []crawler.Document
	if json.Unmarshal(data, &docs) != nil {
		return errors.New("не удалось раскодировать данные кэша")
	}

	s.index.Create(docs)
	s.storage.Create(docs)
	return nil
}

func (s *Service) Search(str string) (docs []crawler.Document) {
	ids := s.index.Find(str)
	for _, id := range ids {
		if doc, ok := s.storage.Document(id); ok != false {
			docs = append(docs, doc)
		}
	}
	return docs
}

// Save - сохраняет документы в кэш
func (s *Service) Save(docs []crawler.Document) error {
	str, err := json.Marshal(docs)
	if err != nil {
		return errors.New("не удалось закодировать данные для кэша")
	}

	return s.cache.Save(str)
}

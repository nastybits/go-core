// Package engine - пакет поискового движка
package engine

import (
	"encoding/json"
	"errors"
	"go.core/lesson6/pkg/cache"
	"go.core/lesson6/pkg/crawler"
	"go.core/lesson6/pkg/index"
	"go.core/lesson6/pkg/storage"
)

type Engine struct {
	Index   index.Service
	Storage storage.Storager
	Cache 	cache.SaveLoader
}

// Load - загружает документы из кэша, индексирует и сохраняет их в хранилище
func (e *Engine) Load() error {
	data, err := e.Cache.Load()
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

	e.Index.Create(docs)
	e.Storage.Create(docs)
	return nil
}

// Save - сохраняет документы в кэш
func (e *Engine) Save(docs []crawler.Document) error {
	str, err := json.Marshal(docs)
	if err != nil {
		return errors.New("не удалось закодировать данные для кэша")
	}

	return e.Cache.Save(str)
}

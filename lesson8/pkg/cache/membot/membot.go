// Package membot - пакет-заглушка для локального кэша
package membot

import (
	"encoding/json"
	"go.core/lesson8/pkg/fixtures"
)

type Cache struct {}

func (c *Cache) Load() ([]byte, error) {
	docs := fixtures.Documents()
	return json.Marshal(&docs)
}

func (c *Cache) Save(p []byte) error {
	return nil
}

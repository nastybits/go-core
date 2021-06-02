// Package membot - пакет-заглушка для локального кэша
package membot

import "encoding/json"

type Cache struct {}

type doc struct {
	ID int
	URL string
	Title string
}

func (c *Cache) Load() ([]byte, error) {
	docs := []doc{
		{ID: 1, URL: "test1.com", Title: "Test 1"},
		{ID: 2, URL: "test2.com", Title: "Test 2"},
		{ID: 3, URL: "test3.com", Title: "Test 3"},
		{ID: 4, URL: "test4.com", Title: "Test 4"},
		{ID: 5, URL: "test5.com", Title: "Test 5"},
	}

	return json.Marshal(&docs)
}

func (c *Cache) Save(p []byte) error {
	return nil
}

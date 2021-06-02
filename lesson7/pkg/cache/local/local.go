// Package local - пакет позволяет рабоать с локальным (файловым) кэшем
package local

import (
	"io/ioutil"
	"os"
)

type Cache struct {
	path string
}

func New(path string) *Cache {
	c := Cache{path: path}
	return &c
}

// Load - загрузить данные из кэша
func (c *Cache) Load() ([]byte, error) {
	_, err := os.Stat(c.path)
	if os.IsNotExist(err) {
		_, err = os.Create(c.path)
		if err != nil {
			return []byte{}, err
		}
	}
	return ioutil.ReadFile(c.path)
}

// Save - сохранить данные в кэш
func (c *Cache) Save(p []byte) error {
	return ioutil.WriteFile(c.path, p,777)
}

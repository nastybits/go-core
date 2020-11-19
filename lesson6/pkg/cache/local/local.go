// Package local - пакет позволяет рабоать с локальным (файловым) кэшем
package local

import (
	"io/ioutil"
	"os"
)

type Cache struct {
	Path string
}

// Load - загрузить данные из кэша
func (c *Cache) Load() ([]byte, error) {
	_, err := os.Stat(c.Path)
	if os.IsNotExist(err) {
		_, err = os.Create(c.Path)
		if err != nil {
			return []byte{}, err
		}
	}
	return ioutil.ReadFile(c.Path)
}

// Save - сохранить данные в кэш
func (c *Cache) Save(p []byte) error {
	return ioutil.WriteFile(c.Path, p,777)
}

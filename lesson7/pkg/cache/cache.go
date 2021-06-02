// Package Cache - пакет определяет контракт для пакетов кэширования
package cache

type Interface interface {
	Load() ([]byte, error)
	Save(p []byte) error
}

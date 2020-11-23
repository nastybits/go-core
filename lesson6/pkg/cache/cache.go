// Package Cache - пакет определяет контракт для пакетов кэширования
package cache

type SaveLoader interface {
	Load() ([]byte, error)
	Save(p []byte) error
}

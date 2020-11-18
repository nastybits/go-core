// Package crawler - пакет сканирует ресурсы с использование сканнера и возвращает структуры
// реализующие контракт document.Documenter
package crawler

import (
	"go.core/lesson4/pkg/document"
	"sort"
)

type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

// Scan - осуществляет сканирование с помощью переданного сканера и возвращает документы реализующие
// интерфейс Documenter в порядке возрастания их ID
func Scan(s Scanner, urls []string, depth int) (documents []document.Documenter, err error) {
	for _, url := range urls {
		data, err := s.Scan(url, depth)
		if err != nil {
			return documents, err
		}
		var id int
		for url, title := range data {
			documents = append(documents, document.New(id, url, title))
			id++
		}
	}
	sort.Slice(documents, func(i, j int) bool { return documents[i].Id() < documents[j].Id() })
	return documents, nil
}

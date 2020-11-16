// Package crawler - пакет сканирует ресурсы с использование сканнера и возвращает структуры
// реализующие контракт document.Documenter
package crawler

import (
	"regexp"
	"sort"
	"strings"
)

type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

type Service struct {
	Scanner Scanner
}

type Document struct {
	ID int
	URL string
	Title string
}

func New(s Scanner) Service {
	return Service{
		Scanner: s,
	}
}

// Scan - осуществляет сканирование с помощью переданного сканера и возвращает документы реализующие
// интерфейс Documenter в порядке возрастания их ID
func (s *Service) Scan(urls []string, depth int) (documents []Document, err error) {
	var id int
	for _, url := range urls {
		data, err := s.Scanner.Scan(url, depth)
		if err != nil {
			return documents, err
		}
		for url, title := range data {
			documents = append(documents, Document{ID: id, URL: url, Title: title})
			id++
		}
	}

	sort.Slice(documents, func(i, j int) bool { return documents[i].ID < documents[j].ID })
	return documents, nil
}

func (d *Document) Words() []string {
	str := strings.Trim(strings.ToLower(d.URL + " " + d.Title),"\r\n ")
	str = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9]`).ReplaceAllString(str, " ")
	str = regexp.MustCompile("\\s+").ReplaceAllString(str, ",")
	return strings.Split(str,",")
}

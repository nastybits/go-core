// Package crawler - пакет сканирует ресурсы с использование сканнера и возвращает структуры
// реализующие контракт document.Documenter
package crawler

import (
	"regexp"
	"sort"
	"strings"
)

type Interface interface {
	Scan(urls []string, depth int) ([]Document, error)
}

// Scanner - интерфейс обеспечивает контракт сканнера для карулера
type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

type Crawler struct {
	Scanner Scanner
}

// Document - структура документов возвращаемых краулером
type Document struct {
	ID int
	URL string
	Title string
}

// New - конструктор краулера
func New(s Scanner) *Crawler {
	c := Crawler{
		Scanner: s,
	}
	return &c
}

// Scan - осуществляет сканирование с помощью переданного сканера и возвращает документы реализующие
// интерфейс Documenter в порядке возрастания их ID
func (c *Crawler) Scan(urls []string, depth int) (documents []Document, err error) {
	var id int
	for _, url := range urls {
		data, err := c.Scanner.Scan(url, depth)
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

// Words - Разбивает содержимое текстовых полей документа по словами возвращает их ввиде среза
func (d *Document) Words() []string {
	str := strings.Trim(strings.ToLower(d.URL + " " + d.Title),"\r\n ")
	str = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9]`).ReplaceAllString(str, " ")
	str = regexp.MustCompile("\\s+").ReplaceAllString(str, ",")
	return strings.Split(str,",")
}

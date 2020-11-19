// Package index - пакет индексирует документы и хранит их обратный индекс
package index

import (
	"go.core/lesson6/pkg/crawler"
)

type Service map[string][]int

func New() Service {
	return make(Service)
}

func clear(s Service) {
	s = map[string][]int{}
}

// Index - создает и сохраняет обратный индекс переданных документов
func (s Service) Create(d []crawler.Document) {
	clear(s)
	for _, doc := range d {
		s.Add(doc)
	}
}

// Add - индексирует переданный документ и сохраняет его обратный индекс в Storage
func (s Service) Add(d crawler.Document) {
	for _, word := range d.Words() {
		items, ok := s[word]
		if !ok {
			s[word] = []int{d.ID}
		}
		if !s.inArray(items, d.ID) {
			s[word] = append(items, d.ID)
		}
	}
}

// Find - поиск ID документа по строке в обратном индексе
func (s Service) Find(str string) (i []int) {
	i, _ = s[str]
	return i
}

// inArray - поиск элемента в срезе. Возвращает его индекс или -1 если элемент не найден
func (s Service) inArray(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

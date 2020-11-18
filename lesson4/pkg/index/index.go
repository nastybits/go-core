// Package index - пакет индексирует документы и хранит их обратный индекс
package index

import (
	"go.core/lesson4/pkg/document"
)

type Index map[string][]int

func New() Index {
	s := make(Index)
	return s
}

func (s Index) Init(d []document.Documenter) {
	for _, doc := range d {
		s.Add(doc)
	}
}

// Index - индексирует переданный документ и сохраняет его обратный индекс в Storage
func (s Index) Add(d document.Documenter) {
	for _, word := range d.Words() {
		items, ok := s[word]
		if !ok {
			s[word] = []int{d.Id()}
		}
		if !inArray(items, d.Id()) {
			s[word] = append(items, d.Id())
		}
	}
}

func (s Index) Find(str string) (i []int) {
	i, _ = s[str]
	return i
}

// inArray - ищет элемент в срезе и возвращает его индекс или -1 если элемент не найден
func inArray(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

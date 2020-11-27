// Package index - пакет индексирует документы и хранит их обратный индекс
package index

import (
	"go.core/lesson8/pkg/crawler"
)

type Interface interface {
	Create(d []crawler.Document)
	Add(d crawler.Document)
	Find(str string) (i []int)
}

type Index struct {
	data map[string][]int
}

func New() *Index {
	s := Index{
		data: make(map[string][]int),
	}
	return &s
}

// Create - создает и сохраняет обратный индекс переданных документов
func (idx *Index) Create(d []crawler.Document) {
	clear(idx)
	for _, doc := range d {
		idx.Add(doc)
	}
}

// Add - индексирует переданный документ и сохраняет его обратный индекс
func (idx *Index) Add(d crawler.Document) {
	for _, word := range d.Words() {
		items, ok := idx.data[word]
		if !ok {
			idx.data[word] = []int{d.ID}
		}
		if !idx.inArray(items, d.ID) {
			idx.data[word] = append(items, d.ID)
		}
	}
}

// Find - поиск ID документа по строке в обратном индексе
func (idx Index) Find(str string) (i []int) {
	i, _ = idx.data[str]
	return i
}

// clear - очищает индекс
func clear(idx *Index) {
	idx.data = make(map[string][]int)
}

// inArray - поиск элемента в срезе. Возвращает его индекс или -1 если элемент не найден
func (idx *Index)inArray(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

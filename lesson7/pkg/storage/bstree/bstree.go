// Package bstree - хранилище документов в виде двоичного дерева поиска.
// Пакет обеспечивает операции вставки, удаления и поиска документа
package bstree

import "go.core/lesson7/pkg/crawler"

// Tree структура для работы с элементами. Представляет из себя двоичное дерево поиска
type Tree struct {
	root *Element
}

// Element структура для хранения элемента дерева
type Element struct {
	left, right *Element
	Doc         crawler.Document
}

// Document - поиск элемента по его ID, возвращает элемент и true если найдено, иначе false
func (t *Tree) Document(id int) (crawler.Document, bool) {
	if el, ok := search(t.root, id); ok != false {
		return el.Doc, true
	}
	return crawler.Document{}, false
}

func (t *Tree) Create(docs []crawler.Document) {
	t.root = nil
	for _, doc := range docs {
		t.Add(doc)
	}
}

// Add - вставляет элемент в дерево
func (t *Tree) Add(doc crawler.Document) {
	el := &Element{Doc: doc}
	if t.root == nil {
		t.root = el
		return
	}
	insert(t.root, el)
}

// search - рекурсивно ищет элемент в дереве
func search(d *Element, id int) (doc *Element, res bool) {
	if d == nil {
		return doc, false
	}
	if d.Doc.ID == id {
		return d, true
	}
	if d.Doc.ID < id {
		return search(d.right, id)
	}
	return search(d.left, id)
}

// insert - рекурсивно вставляет элемент в дерево
func insert(node, new *Element) {
	if new.Doc.ID < node.Doc.ID {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.Doc.ID >= node.Doc.ID {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

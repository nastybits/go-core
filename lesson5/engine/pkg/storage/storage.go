package storage

type Documents struct {
	root *Document
}

type Document struct {
	left, right *Document
	ID          int
	Url         string
	Title       string
}

// Document - поиск документа в дереве, выдаёт true если найдено, иначе false
func (ds *Documents) Document(id int) (*Document, bool) {
	return search(ds.root, id)
}

// Add - вставляет документ в дерево
func (ds *Documents) Add(d Document) {
	if ds.root == nil {
		ds.root = &d
		return
	}
	ds.insert(ds.root, &d)
}

// Delete - вставляет документ в дерево
func (ds *Documents) Delete(id int) error {
	return nil
}

// insert - рекурсивно вставляет документ в дерево
func (ds *Documents) insert(node, new *Document) {
	if new.ID < node.ID {
		if node.left == nil {
			node.left = new
			return
		}
		ds.insert(node.left, new)
	}
	if new.ID >= node.ID {
		if node.right == nil {
			node.right = new
			return
		}
		ds.insert(node.right, new)
	}
}

// search - рекурсивный поиск документа в дереве
func search(d *Document, id int) (doc *Document, res bool) {
	if d == nil {
		return doc, false
	}
	if d.ID == id {
		return d, true
	}
	if d.ID < id {
		return search(d.right, id)
	}
	return search(d.left, id)
}

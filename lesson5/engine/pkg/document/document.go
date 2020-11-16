// Package document - пакет обеспечивает создание структур документа реализующих контракт Documenter
package document

import (
	"fmt"
	"regexp"
	"strings"
)

type Documenter interface {
	Id() int
	Title() string
	Words() []string
	fmt.Stringer
}

type WebPage struct {
	id int
	title string
	url string
}

func New(id int, url string, title string) Documenter {
	return WebPage{id: id, title: title, url: url}
}

// Getter Id
func (d WebPage) Id() int {
	return d.id
}

// Getter Title
func (d WebPage) Title() string {
	return d.title
}

// Getter Url
func (d WebPage) Url() string {
	return d.url
}

func (d WebPage) Words() []string {
	str := strings.Trim(strings.ToLower(d.url + " " + d.title),"\r\n ")
	str = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9]`).ReplaceAllString(str, " ")
	str = regexp.MustCompile("\\s+").ReplaceAllString(str, ",")
	return strings.Split(str,",")
}

func (d WebPage) String() string {
	return fmt.Sprintf("%s: %s", d.url, d.title)
}

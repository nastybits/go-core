// Package main - задание четвертого урока для курса go-core.
package main

import (
	"fmt"
	"go.core/lesson4/engine/pkg/crawler"
	"go.core/lesson4/engine/pkg/crawler/spider"
	"go.core/lesson4/engine/pkg/document"
	"go.core/lesson4/engine/pkg/index"
	"sort"
	"strings"
)

func main() {
	urls := []string{"https://golangs.org", "https://altech.online"}

	fmt.Print("Сканирование сайтов ")
	var s spider.Scanner
	docs, err := crawler.Scan(s, urls, 2)
	if err != nil {
		return
	}
	fmt.Println("завершено.")

	fmt.Print("Индексирование страниц ")
	storage := index.New()
	storage.Init(docs)
	fmt.Println("завершено.")

	var q string
	for {
		fmt.Print("Введите поисковый запрос: ")
		_, err := fmt.Scanf("%s\n", &q)

		if err != nil {
			fmt.Println("Программа завершила работу.")
			return
		}

		IDs := storage.Find(strings.ToLower(q))
		res := find(IDs, docs)
		fmt.Printf("Результаты поиска по запросу \"%s\":\nНайдено всего: %d\n", q, len(res))
		for _, doc := range res {
			fmt.Println(doc)
		}
	}
}

func find(IDs []int, d []document.Documenter) (res []document.Documenter) {
	for _, id := range IDs {
		idx := sort.Search(len(d), func(i int) bool { return d[i].Id() >= id })
		if idx < len(d) {
			res = append(res, d[idx])
		}
	}
	return res
}

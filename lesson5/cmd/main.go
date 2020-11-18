// Package main - задание пятого урока для курса go-core.
package main

import (
	"fmt"
	"go.core/lesson5/pkg/crawler"
	"go.core/lesson5/pkg/crawler/spider"
	"go.core/lesson5/pkg/index"
	"go.core/lesson5/pkg/storage"
	"go.core/lesson5/pkg/storage/bstree"
	"strings"
)

type Engine struct {
	Index   index.Service
	Storage storage.Storager
}

func main() {
	urls := []string{"https://golangs.org", "https://altech.online"}
	var binaryTree bstree.Tree

	// Инициализируем краулер со сканером пауком
	var s spider.Scanner
	c := crawler.New(s)

	// Инициализация движка с сервисами
	engine := Engine{
		Index:   index.New(),
		Storage: storage.New(&binaryTree),
	}

	fmt.Print("Сканирование сайтов ")
	docs, err := c.Scan(urls, 2)
	if err != nil {
		return
	}
	fmt.Println("завершено.")

	fmt.Print("Индексирование страниц ")
	for _, doc := range docs {
		engine.Index.Add(doc)
		engine.Storage.Add(doc)
	}
	fmt.Println("завершено.")

	var str string
	for {
		fmt.Print("Введите поисковый запрос: ")
		_, err := fmt.Scanf("%s\n", &str)

		if err != nil {
			fmt.Println("Программа завершила работу.")
			return
		}

		IDs := engine.Index.Find(strings.ToLower(str))
		var res []crawler.Document
		for _, id := range IDs {
			if d, ok := engine.Storage.Document(id); ok != false {
				res = append(res, d)
			}
		}

		fmt.Printf("Результаты поиска по запросу \"%s\":\nНайдено всего: %d\n", str, len(res))
		for _, doc := range res {
			fmt.Println(doc)
		}
	}
}

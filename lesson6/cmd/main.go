// Package main - задание шестого урока для курса go-core.
package main

import (
	"fmt"
	"go.core/lesson6/pkg/cache/local"
	"go.core/lesson6/pkg/crawler"
	"go.core/lesson6/pkg/crawler/spider"
	"go.core/lesson6/pkg/engine"
	"go.core/lesson6/pkg/index"
	"go.core/lesson6/pkg/storage"
	"go.core/lesson6/pkg/storage/bstree"
	"strings"
)

func main() {
	urls := []string{
		"https://altech.online",
		"https://www.coffeestainstudios.com",
		"https://golangs.org",
		"https://www.cyberpunk.net/ru/ru/",
	}

	var bt bstree.Tree
	cf := local.Cache{Path: "../data/storage.txt"}

	// Инициализация движка с сервисами
	e := engine.Engine{
		Index:   index.New(),
		Storage: storage.New(&bt), // Хранилище для быстрого поиска и вставки (Бинарное дерево)
		Cache:   &cf, // Локальный кэш
	}

	// Пробуем загрузить данные из файла хранилища
	if err := e.Load(); err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем краулер со сканером пауком
	var s spider.Scanner
	c := crawler.New(s)

	// Запускаем сканирование парралельно основной работе
	go func() {
		docs, err := c.Scan(urls, 2)
		if err != nil {
			return
		}

		// Обновляем хранилище в памяти и индекс
		e.Index.Create(docs)
		e.Storage.Create(docs)

		// Сохраняем данные сканера в файл хранилища в формате JSON
		if e.Save(docs) != nil {
			fmt.Println("Не удалось сохранить данные в хранилище")
			return
		}
	}()

	var str string
	for {
		fmt.Print("Введите поисковый запрос: ")
		_, err := fmt.Scanf("%s\n", &str)

		if err != nil {
			fmt.Println("Программа завершила работу.")
			return
		}

		IDs := e.Index.Find(strings.ToLower(str))
		var res []crawler.Document
		for _, id := range IDs {
			if d, ok := e.Storage.Document(id); ok != false {
				res = append(res, d)
			}
		}

		fmt.Printf("Результаты поиска по запросу \"%s\":\nНайдено всего: %d\n", str, len(res))
		for _, doc := range res {
			fmt.Println(doc)
		}
	}
}

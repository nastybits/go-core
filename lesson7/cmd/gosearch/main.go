// Package main - пакет поискового робота для задания 7
package main

import (
	"fmt"
	"go.core/lesson7/pkg/cache/local"
	"go.core/lesson7/pkg/crawler"
	"go.core/lesson7/pkg/crawler/spider"
	"go.core/lesson7/pkg/engine"
	"go.core/lesson7/pkg/index"
	"go.core/lesson7/pkg/storage"
	"strings"
)

// Сервер поисковика GoSearch
type gosearch struct {
	engine  *engine.Service
	scanner crawler.Interface
	index   index.Interface
	storage storage.Interface
}

// new - Конструктор gosearch
func new() *gosearch {
	var s spider.Scanner

	gs := gosearch{}
	gs.scanner 	= crawler.New(s)
	gs.index 	= index.New()
	gs.storage 	= storage.New()
	gs.engine 	= engine.New(gs.index, gs.storage, local.New("../data/storage.txt"))

	return &gs
}

// init - функция инциализации. Запускает сканирование и обрабатыает полученные данные
func (gs *gosearch) init(urls []string) {
	err := gs.engine.Load() // Пробуем загрузить данные из файла хранилища
	if err != nil { // Если не удалось - запускаем сканирование синхронно
		docs, err := gs.scanner.Scan(urls, 2)
		if err != nil {
			fmt.Println("Не удалось просканировать сайты")
			return
		}
		gs.index.Create(docs)
		gs.storage.Create(docs)
		err = gs.engine.Save(docs)
		if err != nil {
			fmt.Println("Не удалось закэшировать данные")
		}
		return
	}

	go func() { // Если удалось - запускаем сканирование асинхронно и продолжаем работу
		docs, err := gs.scanner.Scan(urls, 2)
		if err != nil {
			fmt.Println("Не удалось просканировать сайты")
			return
		}
		gs.index.Create(docs)
		gs.storage.Create(docs)
		err = gs.engine.Save(docs)
		if err != nil {
			fmt.Println("Не удалось закэшировать данные")
		}
	}()
}

// run - функция запускает интерактивный интерфейс командной строки для поиска
func (gs *gosearch) run() {
	var str string
	for {
		fmt.Print("Введите поисковый запрос: ")
		_, err := fmt.Scanf("%s\n", &str)

		if err != nil {
			fmt.Println("Программа завершила работу.")
			return
		}

		docs := gs.engine.Search(strings.ToLower(str))
		fmt.Printf("Результаты поиска по запросу \"%s\":\nНайдено всего: %d\n", str, len(docs))
		for _, doc := range docs {
			fmt.Println(doc)
		}
	}
}

func main() {
	urls := []string{
		"https://altech.online",
		"https://www.coffeestainstudios.com",
		"https://golangs.org",
		"https://www.cyberpunk.net/ru/ru/",
	}

	gs := new()
	gs.init(urls)
	gs.run()
}

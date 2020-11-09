// Package main - задание вторго урока для курса go-core.
package main

import (
	"fmt"
	"go.core/lesson2/search/crawler/pkg/spider"
	"go.core/lesson2/search/engine/pkg/search"
)

func main() {
	sites := [2]string{"https://golangs.org", "https://altech.online"}
	pages := make(map[string]string)

	fmt.Print("Сканирование сайтов ")
	for _, site := range sites {
		data, err := spider.Scan(site, 2)
		if err != nil {
			return
		}

		for url, title := range data {
			pages[url] = title
		}
	}
	fmt.Println("завершено.")

	var text string
	for {
		fmt.Print("Введите поисковый запрос: ")
		_, err := fmt.Scanf("%s\n", &text)

		if err != nil {
			fmt.Print("Программа завершила работу.")
			return
		}

		res := search.Word(text, pages)
		if len(res) == 0 {
			fmt.Printf("По запросу \"%s\" ничего не найдено\n", text)
		}

		if len(res) > 0 {
			fmt.Printf("Результаты поиска по запросу \"%s\":\nНайдено всего: %d\n", text, len(res))
			for url, title := range res {
				fmt.Printf("%s - %s\n", url, title)
			}
		}
	}
}

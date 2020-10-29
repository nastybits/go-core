// Package main - задание вторго урока для курса go-core.
package main

import (
	"education/lesson3/engine/pkg/scanner"
	"fmt"
	"strings"
)

func main() {
	urls := [2]string{"https://golangs.org", "https://altech.online"}
	pages := make(map[string]string)
	var s scanner.Web

	fmt.Print("Сканирование сайтов ")
	for _, url := range urls {
		data, err := scan(s, url, 2)
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
			fmt.Println("Программа завершила работу.")
			return
		}

		res := find(text, pages)
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

// Scanner
type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

// Scan
func scan(s Scanner, url string, depth int) (map[string]string, error) {
	return s.Scan(url, depth)
}

// find поиск подстроки в ключах и значениях ассоциативного массива
// Возвращает ассоциативный массив из элементов переданного массива удовлеворяющих условиям поиска
func find(text string, data map[string]string) (res map[string]string) {
	res = make(map[string]string)
	text = strings.ToLower(text)

	for key, val := range data {
		if strings.Contains(strings.ToLower(val), text) || strings.Contains(strings.ToLower(key), text) {
			res[key] = val
		}
	}

	return res
}

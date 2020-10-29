// Package search - пакет для поиска значений в ассоциативном массиве.
// Пакет позволяет получить ассоциативный массив ключ - значение для найденных элементов
package search

import (
	"strings"
)

// Find поиск подстроки в значениях ассоциативного массива
// Возвращает ассоциативный массив из элементов переданного массива удовлеворяющих условиям поиска
func Word(text string, data map[string]string) (res map[string]string) {
	res = make(map[string]string)
	text = strings.ToLower(text)

	for key, val := range data {
		if strings.Contains(strings.ToLower(val), text) || strings.Contains(strings.ToLower(key), text) {
			res[key] = val
		}
	}

	return res
}

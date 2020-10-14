// Package fibonacci - пакет для получения числа Фибоначчи по его порядковому номеру.
package fibonacci

import "fmt"

// Number - возвращает число Фибоначчи по переданному индексу.
func Number(n int) (res int, err error) {
	if n <= 0 || n > 20 {
		err = fmt.Errorf("Значение должно быть больше 0 и меньше 20")
		return
	}

	return recursive(n), err
}

// Рекурсивный поиск числа Фибоначчи.
func recursive(n int) int {
	if n > 2 {
		return recursive(n-1) + recursive(n-2)
	}
	return n - 1
}

// Динамический поиск числа Фибоначчи.
func dynamic(n int) int {
	a := 0
	b := 1

	if n <= 1 {
		return 0
	}

	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}

	return b
}

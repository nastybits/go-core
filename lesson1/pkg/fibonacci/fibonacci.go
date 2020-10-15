// Package fibonacci - пакет для получения числа Фибоначчи по его порядковому номеру.
package fibonacci

// Number - возвращает число Фибоначчи по переданному индексу.
func Num(n int) int {
	if n > 2 {
		return Num(n-1) + Num(n-2)
	}
	return n - 1
}

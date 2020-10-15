// Package main - задание первого урока для курса go-core.
package main

import (
	"education/lesson1/pkg/fibonacci"
	"flag"
	"fmt"
	"log"
)

func main() {
	n := flag.Int("n", 1, "Индекс числа Фибоначчи")
	flag.Parse()

	if *n <= 0 || *n > 20 {
		log.Fatal("Значение должно быть больше 0 и меньше 20")
	}

	num := fibonacci.Num(*n)

	fmt.Printf("Значение числа Фибоначчи по индексу %d равно: %d", *n, num)
}

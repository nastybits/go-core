// Package main - задание первого урока для курса go-core.
package main

import (
	"flag"
	"fmt"
	"log"

	"education/lesson1/pkg/fibonacci"
)

func main() {
	n := flag.Int("n", 1, "Индекс числа Фибоначчи")
	flag.Parse()

	num, err := fibonacci.Number(*n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Значение числа Фибоначчи по индексу %d равно: %d", *n, num)
}

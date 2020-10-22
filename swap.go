package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a int
	var b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	swap(&a, &b)

	fmt.Println(a, b)
}
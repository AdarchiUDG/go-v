package main

import "fmt"

func fibonacci(n int, current int, last int) {
	fmt.Print(current, " ")
	if n > 0 {
		fibonacci(n - 1, current + last, current)
	}
}

func main() {
	fibonacci(15, 0, 1)
}

package main

import "fmt"

func largestNumber(nums ...int64) int64 {
	var largest int64
	for _, v := range nums {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	fmt.Print(largestNumber(1, 98, 372, 2, 29, 3283, 8, 0))
}
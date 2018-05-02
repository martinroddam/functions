package main

import "fmt"

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	sum(aSlice...)
	sum()
}

func sum(nums ...int) {
	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Total sum of numbers %v is: %v\n", nums, total)
}

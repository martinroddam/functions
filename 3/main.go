package main

import "fmt"

func main() {
	numList := []int{1, 3, 54, 15, 965, 2}
	fmt.Printf("Provided numbers: %v\n", numList)

	greatest := findGreatest(numList...)
	fmt.Printf("The largest number is: %v\n", greatest)
}

func findGreatest(numbers ...int) int {
	var greatest int
	for _, num := range numbers {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}

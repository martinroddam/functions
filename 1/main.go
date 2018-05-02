package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		result, isEven := half(i)
		fmt.Printf("Given num: %v, Divide result: %v, Is even: %v\n", i, result, isEven)
	}
}

func half(num int) (float32, bool) {
	result := float32(num) / 2.0
	isEven := num%2 == 0
	return result, isEven
}

package main

import "fmt"

func main() {
	fmt.Printf("(true && false) || (false && true) || !(false && false) evaluates to: %v\n", (true && false) || (false && true) || !(false && false))
}

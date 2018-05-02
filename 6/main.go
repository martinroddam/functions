package main

import (
	"fmt"
	"math"
)

var sqrt float64

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func main() {

	// Try values of a and b
	for a := 1; a < 1000; a++ {
		// b must be > a
		for b := a + 1; b < (1000 - a); b++ {

			// Get the product of a^2 + b^2
			firstProduct := math.Pow(float64(a), 2) + math.Pow(float64(b), 2)
			fmt.Printf("Trying a: %v, b: %v\n", a, b)

			sqrt = math.Sqrt(firstProduct)

			// Is the square root of that product a whole number?
			// If yes, then c in the equation above is a natural (whole) number and we have ourselves
			// a Pythagorean triplet!
			if (math.Trunc(sqrt) - sqrt) == 0 {
				// sqrt is a natural number and we can use as the value c
				// a, b, c make up our Pythagorean triplet
				c := int(sqrt)
				if a+b+c == 1000 {
					fmt.Printf("\nSolution!\n\na=%v\nb=%v\nc=%v\n", a, b, c)
					return
				}
			}
		}
	}
}

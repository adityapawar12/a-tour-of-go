package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Like for, the if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// Variables declared inside an if short statement are also available inside any of the else blocks.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 5),
		pow(3, 3, 20),
	)
}

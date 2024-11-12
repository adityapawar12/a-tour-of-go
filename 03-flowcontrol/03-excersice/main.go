package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	n := x

	var root float64

	count := 0

	for true {
		count++
		root = 0.5 * (n + (x / n))
		if math.Abs(root-n) < 0.00001 {
			break
		}
		n = root
	}

	return root
}

func main() {
	fmt.Printf("%.4f", Sqrt(2))
	fmt.Println()
	fmt.Printf("%.4f", Sqrt(3))
	fmt.Println()
	fmt.Printf("%.4f", Sqrt(4))
	fmt.Println()
	fmt.Printf("%.4f", Sqrt(5))
	fmt.Println()
	fmt.Printf("%.4f", Sqrt(6))
	fmt.Println()
	fmt.Printf("%.4f", math.Sqrt(2))
	fmt.Println()
	fmt.Printf("%.4f", math.Sqrt(3))
	fmt.Println()
	fmt.Printf("%.4f", math.Sqrt(4))
	fmt.Println()
	fmt.Printf("%.4f", math.Sqrt(5))
	fmt.Println()
	fmt.Printf("%.4f", math.Sqrt(6))
	fmt.Println()
}

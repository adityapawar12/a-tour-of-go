package main

import (
	"fmt"
)

// A simple go function. It can take zero or more arguments.
// The type comes after the variable name.
func add(numOne int, numTwo int) int {
	return numOne + numTwo
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func substract(numOne, numTwo int) int {
	return numOne - numTwo
}

// A function can return any number of results.
func swap(a, b string) (string, string) {
	return b, a
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(8, 3))
	fmt.Println(add(7, 5))
	fmt.Println("")

	fmt.Println(substract(1, 2))
	fmt.Println(substract(8, 3))
	fmt.Println(substract(7, 5))
	fmt.Println("")

	fmt.Println(swap("a", "b"))
	fmt.Println(swap("c", "d"))
	fmt.Println(swap("e", "f"))
	fmt.Println("")

	fmt.Println(split(12))
	fmt.Println(split(83))
	fmt.Println(split(75))
	fmt.Println("")
}

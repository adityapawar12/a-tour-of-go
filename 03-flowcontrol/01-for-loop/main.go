package main

import "fmt"

func main() {
	sum := 0

	// Go has only one looping construct, the for loop.
	// the init statement: executed before the first iteration
	// the condition expression: evaluated before every iteration
	// the post statement: executed at the end of every iteration
	// the variables declared there are visible only in the scope of the for statement.
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Iteration Number ", i, " Value is ", sum)
	}
	fmt.Println("Sum is ", sum)

	newSum := 1

	// The init and post statements are optional. At that point you can drop the semicolons: C's while is spelled for in Go.
	for newSum < 1000 {
		newSum += newSum
		fmt.Println("Sum is ", newSum)
	}

	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed
	// for {
	// 	fmt.Println("test")
	// }
}

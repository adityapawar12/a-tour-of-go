package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.
	// Go only runs the selected case, not all the cases that follow.
	// In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go.
	// Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

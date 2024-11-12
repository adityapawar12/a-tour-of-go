package main

import "fmt"

func main() {
	// Go has pointers. A pointer holds the memory address of a value.
	i, j := 42, 2701

	// The & operator generates a pointer to its operand.
	p := &i // point to i
	// The type *T is a pointer to a T value. Its zero value is nil.
	// The * operator denotes the pointer's underlying value.
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j // point to j
	// This is known as "dereferencing" or "indirecting". Unlike C, Go has no pointer arithmetic.
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

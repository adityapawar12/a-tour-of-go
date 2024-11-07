// This is the name of the package. Every go program starts running from main package.
// The package name is the same as the last element of the import path. ex: the "math/rand" package comprises files that begin with the statement package rand
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favourite number is >>> ", rand.Intn(10))
}

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// The var statement declares a list of variables; as in function argument lists, the type is last.
// Package level var statement.
var c, python, java bool

// A var declaration can include initializers, one per variable.
var rust, cpp = true, true

// Constants are declared like variables, but with the const keyword. They can be character, string, boolean, or numeric values.
const Pi = 3.14

// Numeric constants are high-precision values. An untyped constant takes the type needed by its context.
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// Function level var statement.
	var i int
	fmt.Println(i, c, python, java)

	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var x, y = 6, 9
	fmt.Println(x, y, rust, cpp)

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	k := 3
	fmt.Println(k)

	// Variable declarations may be "factored" into blocks, as with import statements.
	// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables declared without an explicit initial value are given their zero value.
	var p int
	var q float64
	var r bool
	var s string
	fmt.Printf("%v %v %v %q\n", p, q, r, s)

	// The expression T(v) converts the value v to the type T.
	// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	var t1, t2 int = 3, 4

	var flt float64 = math.Sqrt(float64(t1*t1 + t2*t2))
	var uns uint = uint(flt)
	fmt.Println(t1, t2, flt, uns)

	// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.
	v1 := 42
	v2 := 42.33
	v3 := 0.867 + 0.5i
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("v2 is of type %T\n", v2)
	fmt.Printf("v3 is of type %T\n", v3)

	// Constants cannot be declared using the := syntax.
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

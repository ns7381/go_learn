package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	ToBe   bool
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14
func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	//change type
	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(fl)
	fmt.Println(x, y, z)

	v := 42.31232112 // change me!
	fmt.Printf("v is of type %T\n", v)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
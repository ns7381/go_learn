package main

import (
	"fmt"
)

func sum(x, y int) (int) {
	return y + x
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func compute(f func(int, int) int) int {
	return f(3, 4)
	
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Point struct {
	x, y int
}

func (p Point) cal() int {
	p.x = 3
	return p.x * p.y
}

type v int

func (x v) calculate() v {
	return x + 1;
}

func main() {
	//fmt.Println(sum(1,2))
	//fmt.Println(swap("hello", "world"))
	//fmt.Println(split(17))
	//f := func(i, j int) int {
	//	return i + j
	//}
	//fmt.Println(compute(f))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	point := Point{2, 3}
	fmt.Println(point.cal())
	fmt.Println(point.x)
	fmt.Println(v(123).calculate())


}

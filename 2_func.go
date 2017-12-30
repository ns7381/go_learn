package main

import "fmt"

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

func main() {
	fmt.Println(sum(1,2))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
}

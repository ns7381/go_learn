package main

import "fmt"

func main() {
	i, j := 12, 23
	p := &i
	var q *int
	q = &j
	fmt.Println(*p, *q)
}
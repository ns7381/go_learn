package main

import "fmt"

var pow1 = []int{1, 2, 4, 8, 16}

func main()  {
	for i, v := range pow1 {
		fmt.Printf("2^%d=%d\n", i, v)
	}
	p := make([]int, 10)
	for i := range p {
		p[i] = 1 << uint(i)
	}
	for _, v := range p {
		fmt.Println(v)
	}
}

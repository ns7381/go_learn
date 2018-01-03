package main

import "fmt"

func main() {
	var a [2]int
	a[1] = 1
	a[0] = 0
	fmt.Println(a)

	s := []int{1, 2, 3, 4}
	for e := range s {
		fmt.Println(s[e])
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]=%d\n", i, s[i])
	}
	ss := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
	}
	fmt.Println(ss[0][1])
	fmt.Println(s[1:3])
	fmt.Println(s[1:])
	fmt.Println(s[:3])

	m1 := make([]int, 5)
	m2 := make([]int, 0, 5)
	m3 := m2[1:3]
	fmt.Println(len(m1))
	fmt.Println(cap(m1))
	fmt.Println(len(m2))
	fmt.Println(cap(m2))
	fmt.Println(len(m3))
	fmt.Println(cap(m3))
	m1 = append(m1, 0)
	fmt.Println(m1)
}

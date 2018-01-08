package main

import "fmt"

type Flyable interface {
	Fly() string
}

type Dunck string

func (d Dunck) Fly() string {
	return string(d)
}

func main() {
	test := Dunck("test")
	var f Flyable
	f = test
	fmt.Println(f.Fly())
	
}

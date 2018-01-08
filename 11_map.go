package main

import "fmt"

type Cell struct {
	width  int
	height int
}

var m map[string]Cell

func main() {
	m2 := map[string]Cell{
		"m1sd": {1, 2},
		"m2":   {1, 2},
	}
	m = make(map[string]Cell)
	m["Test"] = Cell{1, 3}
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m2["m2"])
	delete(m2, "m2")
	fmt.Println(m2)
	v, ok := m2["m2"]
	fmt.Println(v, ok)

}

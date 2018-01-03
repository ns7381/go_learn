package main

import "fmt"

type Student struct {
	name string
	class string
}

var (
	s1 = Student{"na", "1"}
	s2 = Student{name: "b"}
	s3 = Student{}
	s4 = &Student{"d", "4"}
)

func main() {
	student := Student{"nathan", "3.2"}
	fmt.Println(student)
	fmt.Println(student.name)
	s := &student
	s.class = "4.3"
	fmt.Println((*s).name)
	fmt.Println(s.class)
	fmt.Println(s1, s2, s3, s4)
}
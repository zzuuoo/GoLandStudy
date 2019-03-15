package main

import "fmt"

type struct01 struct {
	name string
	age int
}
type struct02 struct {
	struct01
	number string
}

func main() {
	s1 := &struct01{name:"zuo",age:1}
	s := &struct02{number:"123456",struct01:*s1}
	s2 := new(struct02)
	s2.age=19
	s2.name="s2"
	s2.number="s2number"
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
}
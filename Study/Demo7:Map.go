package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "333"
	fmt.Println(m)
	m[6] = "222222"
	fmt.Println(len(m))
	delete(m,1)
	fmt.Println(m)
	v,ok := m[123]
	if ok {
		fmt.Println(v)
	}
	fmt.Println(ok,v)
}

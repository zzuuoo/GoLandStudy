package main

import (
	"fmt"
)

func main() {
	a := [6]int{1,2,3,4,5,6}
	for i,n := range  a{
		fmt.Println(i,n)
	}
	m := make(map[int]string)
	m[9] = "1111"
	m[98] = "hhahhhh"
	for k,v := range  m{
		fmt.Println(k,v)
	}
	que := make(chan string,3)
	que <- "1"
	que <- "2"
	//que <- "3"
	go func() {
		que <- "4"
		close(que)
	}()
	for item := range que {
		fmt.Println(item)
	}
}

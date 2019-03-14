package main

import (
	"fmt"
)

func main() {
	s := make([]string,5)
	fmt.Println(s)
	s[1]= "1111"
	fmt.Println(s)
	s = append(s,"123")
	fmt.Println(s)
	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("c:",c)
	l :=s[1:3]
	fmt.Println(l)
	l = s[:3]
	fmt.Println(l)
	l = s[4:]
	fmt.Println(l)
	fmt.Println(l[0])


	var k []int
	if k == nil{
		fmt.Println("切片为空")
	}
	k = append(k,1)//这样就可以
	fmt.Println(k)

	//k[0] = 1 这样不行，因为切片没申请内存，只是声明

	}
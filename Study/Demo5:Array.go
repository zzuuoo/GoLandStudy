package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)
	arr[4] = 1
	fmt.Println(arr)
	fmt.Println(arr[4])
	a := [2]int{1,2}
	fmt.Println(a)
	var b []int
	b = append(b,1)
	fmt.Println(b)

}

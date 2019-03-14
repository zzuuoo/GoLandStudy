package main

import "fmt"

func main() {
	var k = 1
	switch k {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("default")

	}
	switch  {
	case k>4:
		fmt.Println("k>4")
	case k<2:
		fmt.Println("k<2")
	default:
		fmt.Println("k=<4 and k>=2")

	}
}

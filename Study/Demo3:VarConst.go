package main

import (
	"fmt"
	"math"
)

const name string  = "name"

func main()  {
	fmt.Println(name)
	const a  =  500000000
	const d  =  3e20 / a
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))


	
}

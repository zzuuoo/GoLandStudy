package main

import "fmt"

func main()  {
	for i:=0;i<10;i++{
		fmt.Println("Hello World")
	}
	var k  = 4
	for k < 10 {
		fmt.Println(k)
		k++
	}

	for{
		if k>15{
			break;
		}
		k++
	}
}


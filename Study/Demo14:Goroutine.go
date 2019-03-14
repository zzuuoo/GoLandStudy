package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("开始跑协程")
	go func() {
		defer wg.Done()
		for i:=0;i<3;i++{
			for x := 'a';x < 'a'+26;x++{
				fmt.Printf("%c",x)
			}
		}
		fmt.Println("char")
	}()
	//fmt.Println()

	go func() {
		defer wg.Done()
		for i:=0;i<3;i++{
			for j:=0;j<10;j++{
				fmt.Printf("%d",j)
			}
		}
		fmt.Println("num")
	}()
	fmt.Println()
	fmt.Println("等待协程完成")
	wg.Wait()
	fmt.Println("finished")
}

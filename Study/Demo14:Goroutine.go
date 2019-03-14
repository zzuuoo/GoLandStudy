package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("开始跑协程")
	go func() {
		defer wg.Done()
		for i:=0;i<3;i++{
			for x := 'a';x < 'a'+26;x++{
				fmt.Printf("%c ",x)
			}
		}
		//fmt.Println("char")
	}()
	//fmt.Println()

	go func() {
		defer wg.Done()
		for i:=0;i<3;i++{
			for j:='A';j<'A'+26;j++{
				fmt.Printf("%c ",j)
			}
		}
		//fmt.Println("num")
	}()
	fmt.Println("等待协程完成")
	wg.Wait()
	fmt.Println()
	fmt.Println("finished")
}

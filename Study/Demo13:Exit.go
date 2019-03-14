package main

import (
	"fmt"
	"os"
)
/*
	采用exit 退出程序不会执行 defer 里的内容
 */

func main() {
	defer fmt.Println("我要退出前执行")
	
	os.Exit(2)
}

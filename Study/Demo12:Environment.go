package main

import (
	"fmt"
	"os"
	"strings"
)
/*
	读取环境变量
 */

func main() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0],":",pair[1])
	}
}

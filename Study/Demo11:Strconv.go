package main

import (
	"fmt"
	"strconv"
)
/*
	字符串格式控制转化
 */

func main() {
	//ParseFloat转换成浮点数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	//ParseInt转换成整数
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	//ParseInt会自动识别16进制字符串
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	//转换成无符号整数
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//转换10进制整数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	//返回错误
	//_, e := strconv.Atoi("wat")
	//fmt.Println(e)

	fmt.Println("11:",string(11))
	var ss string
	ss = strconv.Itoa(999)
	fmt.Println("ss:",ss)

}

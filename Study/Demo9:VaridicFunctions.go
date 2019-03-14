package main

import "fmt"
// 参数类型为接口或者error 的时候可以传nil ，其他不可以 error本质也是接口
func sum(nums...int){
	fmt.Println(nums)
	fmt.Println(nums[0])
	var s int
	for n := range nums{
		s = s + n
	}
	fmt.Println(s)
}
type demo struct {
	Name string
}
type demo02 struct {
	D demo
}
func (d demo)Test1(){

}
func test(a interface{}, err error,k int){
	fmt.Println(a)
}
func main() {
	sum(1,2,4,5,6)
	k := []int{1,2,3}
	sum(k...)
	test(nil,nil,1)
	//d2 := new (demo02)
	d := new (demo)
	d.Name = "1"
	//d2.D.Name

}
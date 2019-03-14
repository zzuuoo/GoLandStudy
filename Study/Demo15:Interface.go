package main

import (
	"fmt"
	"math"
)

//接口定义1111
type geometry interface {
	area() float64
	perim() float64
}

//在rect、circle结构体上实现接口方法
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//rect实现了接口的所有方法

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//circle实现了接口的所有方法

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//接口类型作为函数的参数

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	//r，c都实现了接口的所有方法，所以r和c的实例都可以作为接口的实例。
	measure(r)
	measure(c)
}

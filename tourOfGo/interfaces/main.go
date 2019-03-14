package main

import (
	"fmt"
	"math"
)

type Adser interface {
	Abs() float64
}

func main() {
	var a Adser // a 为接口类型
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f // a Myfloat implements Abs
	a = &v // a *Vertex implements Abs

	a = v // a vertex did not implement Abs, 因为没有实现对应的方法，所以类型不匹配，不能赋值

	fmt.Println(a.Abs())
}

type MyFloat float64

// implement Abs method to MyFloat type with value
// Golang没有Implement关键词，是隐式地实现
// 隐式实现分离了接口定义与它的实现
func (f MyFloat) Abs() float64  {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// implement Abs method to Vertex type with pointer receiver
func (v *Vertex) Abs() float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

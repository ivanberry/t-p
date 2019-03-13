package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X, Y int
}

var m map[string]Vertex

var (
	// struct literal denotes a newly allocated struct value
	p1 = Vertex{1,2}
	p2 = Vertex{X: 1}
	p3 = Vertex{}
	pp = &Vertex{1,2}
)

func main() {
	// declare a variable
	var i = 0

	// declare a pointer which point to value with type int
	var n *int
	var x *Vertex
	var x0 *Vertex

	// declare a pointer to i
	var p = &i

	fmt.Printf("n的类型为: %t, %v\n", reflect.TypeOf(n), n)
	fmt.Printf("p的类型为: %t, %v\n", reflect.TypeOf(p), p)
	fmt.Printf("i的类型为: %t, %v\n", reflect.TypeOf(i), &i)
	fmt.Printf("p的内存地址值为: %v\n", &p)

	fmt.Println(p1, p2, p3, pp, x)
	fmt.Printf("type有对应的内存地址吗, %v, %v, %v, %v", &pp, p3, &x, &x0)


	// array and slice
	aInt := []int{1,2,3}

	fmt.Printf("数字类型数组: %v\n", aInt)
	fmt.Println(aInt[0])

	aInt = append(aInt, 12)

	fmt.Println(aInt)

	aIntSlice := aInt[2:]
	fmt.Println(aIntSlice)

	// slice does not store any data, it just describes a section of an underlying array
	aIntSlice[0] = 156
	fmt.Println(aIntSlice, aInt, &aIntSlice[1], &aInt[1], )

	// array/slice literate
	a0 := [2]int{0,1}
	s0 := []int{0, 1}
	fmt.Println(a0, s0, &a0[0], &a0[1])

	// slice length and capacity

	// append to a slice and big than the origin array
	s1 := append(s0, 2, 3, 4)
	fmt.Println(a0, s1, &a0[0], &s0[0], &s1[0], cap(a0), cap(s0), cap(s1))

	// range with for to loop

	for i, v := range s1 {
		v = 10
		fmt.Printf("The %d in slice s1 is %v\n", i, v)
	}
	fmt.Println(s0, s1)

	// make map
	m = make(map[string]Vertex)
	m["ShenZhen"] = Vertex{120, 24}
	m["YuanJiang"] = Vertex{12, 220}

	fmt.Println(m["ShenZhen"])
}
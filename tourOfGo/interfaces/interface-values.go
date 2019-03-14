package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// implement M method with *T type
func (t *T) M()  {
	fmt.Println(t.S)
}

type F float64

// implement M method with F type
func (f F) M()  {
	fmt.Println(f)
}

func main() {
	var i I // declare i of type I, which has method M

	i = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


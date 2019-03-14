package main

import "fmt"

type Vertext struct {
	X, Y float64
}

func (v *Vertext) Scale(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertext, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertext{3,4}
	v.Scale(10)
	(&v).Scale(10)

	p := &Vertext{5, 12}
	ScaleFunc(p, 10)

	fmt.Println(v, p)

}

package main

import (
	"fmt"
	"math"
	"reflect"
)
func Sqrt(x float64) float64 {
	z := x/2
	preZ := 0.0
	for math.Abs(z - preZ) > 0.0000001 {
		preZ = z
		z -= (z*z - x) / (2*z)
		fmt.Printf("current loose: %v %v\n", z, z - preZ)
	}
	return z
}

func main() {
	sum := Sqrt(10)
	fmt.Println(sum)

	var v = 0

	defer fmt.Printf("defer value\n %v", float64(v) )

	fmt.Println(reflect.TypeOf(v))

	defer fmt.Printf("another defer")
	defer fmt.Printf("another defer 0000")

}
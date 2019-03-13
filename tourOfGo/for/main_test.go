package main

import (
	"math"
	"testing"
)

func BenchmarkSqrt(b *testing.B) {
	Sqrt(1000)
}

func BenchmarkSqrt2(b *testing.B) {
	math.Sqrt(1000)
}

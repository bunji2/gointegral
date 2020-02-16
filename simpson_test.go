package main

import (
	"fmt"
	"math"
	"testing"
)

func ExampleSimpson() {
	// interval = [0, 1]
	// n = 1000
	// f(x) = 2*x
	a := float64(0)
	b := float64(1)
	n := int64(1000)
	result := simpson(a, b, n, func(x float64) float64 {
		return float64(2) * x
	})
	fmt.Println("result =", result)
}

func TestSimpson(t *testing.T) {
	a := float64(0.0)
	b := float64(1.0)
	n := int64(1000)
	f := func(x float64) float64 {
		return float64(2.0) * x
	}
	got := simpson(a, b, n, f)
	want := float64(1.0)
	e := float64(0.001) // 0.1[%]
	if math.Abs(got-want)/want > e {
		t.Errorf("got %v want %v", got, want)
	}
}

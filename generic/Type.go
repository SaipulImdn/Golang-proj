package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func SumNumber3[K comparable, V Number](m map[K]V) V {
	var s V
	for _, V := range m {
		s += V
	}
	return s
}

func main() {
	ints := map[string]int64{"first": 34, "seconds": 12}
	floats := map[string]float64{"first": 35.98, "seconds": 26.99}

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumber3(ints),
		SumNumber3(floats))
}

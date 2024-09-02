package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non generic sums int: %v , floact: %v \n", SumInt(ints), SumFloat(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats(floats)) //can call without the type arguments

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

}

func SumInt(m map[string]int64) int64 {
	var sum int64
	for _, key := range m {
		sum += key
	}
	return sum
}

func SumFloat(m map[string]float64) float64 {
	var sum float64
	for _, key := range m {
		sum += key
	}
	return sum
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

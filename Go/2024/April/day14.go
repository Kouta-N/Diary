package main

import (
	"fmt"

	"constraints"
)

type T interface {
	constraints.Ordered
}

func MinSlice[T comparable](slice []T) T {
	min := slice[0]
	for _, val := range slice {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9}
	fmt.Println("Minimum integer:", MinSlice(intSlice))

	floatSlice := []float64{3.14, 2.71, 1.618}
	fmt.Println("Minimum float:", MinSlice(floatSlice))

	stringSlice := []string{"apple", "banana", "orange"}
	fmt.Println("Minimum string:", MinSlice(stringSlice))
}

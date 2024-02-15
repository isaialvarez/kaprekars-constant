package main

import (
	"log"
	"math"
	"slices"
)

func main() {
	var number = 2111

	if number < 100 {
		log.Println("Value must be greater that 99")
		return
	}

	for i := 0; i < 7; i++ {
		array := int2array(number)

		smallest := slices.Clone(array)
		slices.Sort(smallest)

		largest := slices.Clone(smallest)
		slices.Reverse(largest)

		intSmallest := array2Int(smallest)
		intLargest := array2Int(largest)

		if intSmallest == intLargest {

		}

		number = intLargest - intSmallest

		log.Println("Array:", array, "Smallest:", smallest, intSmallest, "Largest:", largest, intLargest, "newNumber", number)

		if number == 6174 || number == 495 {
			break
		}
	}
}

func int2array(value int) []int {
	var array []int

	for value > 0 {
		residual := value % 10
		value /= 10

		array = append(array, residual)
	}

	slices.Reverse(array)

	return array
}

func array2Int(array []int) int {
	result := 0

	for i, v := range array {
		exponential := math.Pow10(len(array) - (i + 1))
		result += v * int(exponential)
	}

	return result
}

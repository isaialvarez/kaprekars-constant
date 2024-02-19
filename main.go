package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Enter number:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("input is not a number")
		return
	}

	generateConstant(number)
}

func generateConstant(number int) int {
	if number < 100 || number > 9998 {
		log.Println("Value must be between 100 and 9998")
		return 0
	}

	fmt.Println("Steps:")
	constant := 0
	for i := 0; i < 7; i++ {
		array := int2array(number)

		smallest := slices.Clone(array)
		slices.Sort(smallest)

		largest := slices.Clone(smallest)
		slices.Reverse(largest)

		intSmallest := array2Int(smallest)
		intLargest := array2Int(largest)

		if intSmallest == intLargest {
			intLargest *= 10
		}

		number = intLargest - intSmallest

		log.Println(intLargest, "-", intSmallest, "=", number)

		constant = number

		if number == 6174 || number == 495 {
			break
		}
	}

	return constant
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

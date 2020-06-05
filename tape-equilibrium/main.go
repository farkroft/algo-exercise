package main

import (
	"fmt"
)

func main() {
	// input := []int{3, 1, 2, 4, 3}
	// input := []int{3, 1}
	// input := []int{-3, -1, -2, -4, -3}
	// input := []int{-1000, 1000}
	input := []int{1, 2, 3, 4, 2} // 0
	output := tapeEquilibrium(input)
	fmt.Println(output)
}

func tapeEquilibrium(arr []int) int {
	var result, totalSum, currentSum int
	arrSum := []int{}

	for i := range arr {
		totalSum += arr[i]
		arrSum = append(arrSum, totalSum)

		if arr[i] < 0 {
			result = result + ^arr[i] + 1
		} else {
			result = result + arr[i]
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		currentSum += arr[i]
		currentResult := currentSum - (totalSum - currentSum)
		if currentResult < 0 {
			currentResult = ^currentResult + 1
		}

		if currentResult < result {
			result = currentResult
		}
	}

	return result
}

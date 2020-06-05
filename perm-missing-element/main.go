package main

import (
	"fmt"
)

func main() {
	input := []int{2, 3, 1, 5} // 4
	output := missingPermElement(input)
	fmt.Println(output)
}

func missingPermElement(arr []int) int {
	sum := 0
	maxElement := len(arr) + 1

	for i := range arr {
		sum += arr[i]
	}

	return maxElement*(maxElement+1)/2 - sum
}

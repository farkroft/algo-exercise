package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	input := []int{9, 3, 9, 3, 9, 7, 9}
	// input := []int{42}
	result := oddOccurance(input)
	fmt.Println(result)

	elapsed := time.Since(start).Seconds()
	fmt.Printf("Binomial took %f", elapsed)
}

// func oddOccurance(arr []int) int {
// 	result := 0
// 	done := false
// 	for !done {
// 		temp := arr[0]
// 		arr = removeElement(arr, 0)
// 		done = true
// 		for i, v := range arr {
// 			if temp == v {
// 				arr = removeElement(arr, i)
// 				done = false
// 				temp = 0
// 			}
// 		}

// 		if temp != 0 {
// 			result = temp
// 		}
// 	}

// 	return result
// }

func removeElement(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func oddOccurance(arr []int) int {
	sort.Ints(arr)
	result := 0
	for i := range arr {
		fmt.Printf("result %d XOR index %d\n", result, arr[i])
		result = result ^ arr[i]
		fmt.Printf("result = %d\n", result)
	}
	return result
}

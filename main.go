package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 2, 4, 6, 7}
	// 7, 1, 3, 2, 4, 5, 6
	// 4, 3, 1, 2
	// 2, 3, 4, 1, 5
	// 1, 3, 5, 2, 4, 6, 7
	swapped := minimumSwap(arr)
	fmt.Println(swapped)
}

// i   arr                         swap (indices)
// 0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
// 1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
// 2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
// 3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
// 4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
// 5   [1, 2, 3, 4, 5, 6, 7]
func minimumSwap(arr []int) int {
	swapTimes := 0
	_, maxNumber := max(arr)

	for i := 1; i <= maxNumber; i++ {
		for j := range arr {
			if i != j+1 {
				if i == arr[j] {
					arr[j], arr[i-1] = arr[i-1], arr[j]
					swapTimes++
				}
			}
		}
	}

	return swapTimes
}

func max(arr []int) (int, int) {
	maxNumber := 0
	position := 0

	for i := range arr {
		if arr[i] > maxNumber {
			maxNumber = arr[i]
			position = i
		}
	}

	return position, maxNumber
}

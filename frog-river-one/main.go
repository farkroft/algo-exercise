package main

import "fmt"

func main() {
	input := []int{2, 2, 2, 2, 2}
	destination := 2
	// 2, [2, 2, 2, 2, 2] // -1
	output := earliestTime(input, destination)
	fmt.Println(output)
}

func earliestTime(arr []int, dest int) int {
	result := false
	timeArrive := 0
	output := 0
	for i := range arr {
		if arr[i] == dest {
			result = true
			break
		} else {
			timeArrive++
		}
	}

	if result == true {
		output = timeArrive
	} else {
		output = -1
	}

	if output == dest || output == 0 {
		output = -1
	}
	return output
}

package main

import "fmt"

func main() {
	input := []int{1, 1, 2, 3, 5}
	rotation := 42
	// input := []int{3, 8, 9, 7, 6}
	// rotation := 20
	// input := []int{3}
	// rotation := 3
	result := cyclicRotation(input, rotation)
	fmt.Println(result)
}

// gagal di K >= N
// func cyclicRotation(arr []int, rotation int) []int {
// 	positions := positionMapping(arr, rotation)
// 	if len(positions) > 1 {
// 		for k, v := range positions {
// 			arr[v] = k
// 		}
// 	}

// 	return arr
// }

// gagal di K >= N
// func positionMapping(arr []int, rotationValue int) map[int]int {
// 	length := len(arr)
// 	result := make(map[int]int)
// 	for i, v := range arr {
// 		pos := i + rotationValue
// 		if pos >= length {
// 			pos = pos % length
// 		}
// 		result[v] = pos
// 	}
// 	return result
// }

func cyclicRotation(arr []int, rotation int) []int {
	positions := positionMapping(arr, rotation)
	result := append([]int(nil), arr...)
	for k := range arr {
		pos := positions[k]
		result[pos] = arr[k]
	}

	return result
}

func positionMapping(arr []int, rotationValue int) []int {
	length := len(arr)
	result := []int{}
	for i := range arr {
		pos := i + rotationValue
		if pos >= length {
			pos = pos % length
		}
		result = append(result, pos)
	}
	return result
}

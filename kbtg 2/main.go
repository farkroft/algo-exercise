package main

import "fmt"

func main() {
	input := [][]int{[]int{2, 3, 2, 0, 1}, []int{7, 1, 1, 2, 6}, []int{5, 1, -7, 1, 8}}
	// fmt.Println(input)
	output := solution(input)
	fmt.Println(output)
}

func solution(arr [][]int) int {
	result := 0
	tempArr := []int{}

	for i := 0; i < len(arr); i++ {
		rowSum := 0
		for j := 0; j < len(arr[i]); j++ {
			rowSum += arr[i][j]
		}
		tempArr = append(tempArr, rowSum)
	}

	for i := 0; i < len(arr); i++ {
		rowSum := 0
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
			rowSum += arr[i][j]
		}
		tempArr = append(tempArr, rowSum)
	}
	// fmt.Println("==========")
	// fmt.Println(tempArr)

	return result
}

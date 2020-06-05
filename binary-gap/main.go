package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 64
	strBinaryInput := strconv.FormatInt(int64(input), 2)
	res := BinaryGap([]rune(strBinaryInput))
	fmt.Println(res)

}

func BinaryGap(input []rune) int {
	var arrBinaryGap []int
	binaryGap := 0
	for i := range input {
		s := fmt.Sprintf("%c", input[i])
		if s == "0" {
			binaryGap++
		}

		if s == "1" {
			arrBinaryGap = append(arrBinaryGap, binaryGap)
			binaryGap = 0
		}
	}

	result := max(arrBinaryGap)
	return result
}

func max(arr []int) int {
	maxValue := arr[0]
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// input := -121 // false
	// input := 121 // true
	input := 1 // true
	res := solution(input)
	fmt.Println(res)
}

func solution(x int) bool {
	arr := []int{}
	arrOfX := intToSlice(x, arr)
	fmt.Println(arrOfX)
	reverseInt := joinIntSlice(arrOfX)
	fmt.Println(reverseInt)
	if x == reverseInt {
		return true
	}
	return false
}

func intToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append(sequence, i)
		return intToSlice(n/10, sequence)
	}

	if n < 0 && n != 0 {
		i := n % 10
		sequence = append(sequence, i)
		return intToSlice(n/-10, sequence)
	}
	return sequence
}

func joinIntSlice(arr []int) int {
	res := 0
	temp := 0

	if len(arr) >= 1 {
		temp = arr[0]
		fmt.Println(temp)
	}
	pow := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		tempPow := math.Pow(float64(10), float64(pow))
		res = res + (arr[i] * int(tempPow))
		pow--
	}

	if temp > 0 {
		return res
	}

	return res * -1
}

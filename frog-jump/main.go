package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10
	y := 85
	d := 30
	result := minimumJumps(x, y, d)
	fmt.Println(result)
}

func minimumJumps(start, finish, leap int) int {
	// var result float64
	// distance := finish - start

	// if start < finish {
	// 	for distance < finish {
	// 		if result == 0 {
	// 			distance = start + leap
	// 			result++
	// 		} else {
	// 			distance = distance + leap
	// 			result++
	// 		}
	// 	}
	// }
	// result = math.Round(float64(distance) / float64(leap))
	return int(math.Ceil((float64(finish - start)) / float64(leap)))

	// return result
}

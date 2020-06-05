package main

import "fmt"

func main() {
	input := []int{2, 1, 1, 10, 2, 13, 3, -18}
	// 2, -3, 3, 1, 10, 8, 2, 5, 13, -5, 3, 18
	// -3, -14, -5, 7, 8, 42, 8, 3
	output := solution(input)
	fmt.Println(output)
	// array represent a year
	// season is array length/4
}

func solution(arr []int) string {
	result := ""
	amplitude := 0
	breakPoint := (len(arr) / 4) - 1
	temp := 0
	season := 1

	for i := range arr {
		temp = ^temp
		temp = temp - arr[i]
		fmt.Printf("%d\t", arr[i])
		if i == breakPoint {
			fmt.Println(temp)
			// fmt.Println(season)
			if temp < 0 {
				temp = ^temp + 1
			}
			// fmt.Println(temp)
			// fmt.Printf("temp %d\n", temp)
			if temp > amplitude {
				amplitude = temp
				if season == 1 {
					result = "WINTER"
				} else if season == 2 {
					result = "SPRING"
				} else if season == 3 {
					result = "SUMMER"
				} else if season == 4 {
					result = "AUTUMN"
				}
			}

			temp = 0
			breakPoint = breakPoint + (len(arr) / 4)

			season++

		}

	}

	// fmt.Println(amplitude)

	return result
}

package main

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := "Sun 10:00-20:00\nFri 05:00-10:00\nFri 16:30-23:50\nSat 10:00-24:00\nSun 01:00-04:00\nSat 02:00-06:00\nTue 03:30-18:15\nTue 19:00-20:00\nWed 04:25-15:14\nWed 15:14-22:40\nThu 00:00-23:59\nMon 05:00-13:00\nMon 15:00-21:00"
	// "Sun 10:00-20:00 Fri 05:00-10:00 Fri 16:30-23:50 Sat 10:00-24:00 Sun 01:00-04:00 Sat 02:00-06:00 Tue 03:30-18:15 Tue 19:00-20:00 Wed 04:25-15:14 Wed 15:14-22:40 Thu 00:00-23:59 Mon 05:00-13:00 Mon 15:00-21:00" // 505 Tue 20:00 - Wed 04:25
	// "Mon 01:00-23:00 Tue 01:00-23:00 Wed 01:00-23:00 Thu 01:00-23:00 Fri 01:00-23:00 Sat 01:00-23:00 Sun 01:00-21:00" // 180 Sun 21:00 - 24:00
	output := solution(input)
	if output > 0 {
		fmt.Println(output)
	}
}

type mapped struct {
	Sun []int
	Mon []int
	Tue []int
	Wed []int
	Thu []int
	Fri []int
	Sat []int
}

func solution(input string) int {
	result := 0
	inputSplitted := strings.Split(input, "\n")
	scheduleMapped := scheduleMappedByDays(inputSplitted)

	valueOfScheduleMapped := reflect.ValueOf(scheduleMapped)

	values := make([]interface{}, valueOfScheduleMapped.NumField())
	for i := range values {
		values[i] = valueOfScheduleMapped.Field(i).Interface()
		sortedArrSchedule := sortedScheduleOneDay(values[i].([]int))
		values[i] = sortedArrSchedule
		fullDaySchedule := weekCompleteSchedule(values[i].([]int))
		fmt.Println(fullDaySchedule)
	}

	return result
}

func convertToInt(input string) {

}

func scheduleMappedByDays(input []string) mapped {
	var scheduleMapped mapped

	for i := range input {
		indexSplitter := strings.Split(input[i], " ")

		switch indexSplitter[0] { // Sun
		case "Sun":
			scheduleSplitted := strings.Split(indexSplitter[1], "-") // 10:00-20:00
			arrOfScheduleInt := strScheduleToInt(scheduleSplitted)

			for j := range arrOfScheduleInt {
				scheduleMapped.Sun = append(scheduleMapped.Sun, arrOfScheduleInt[j])
			}
		case "Mon":
			scheduleSplitted := strings.Split(indexSplitter[1], "-") // 10:00-20:00
			arrOfScheduleInt := strScheduleToInt(scheduleSplitted)

			for j := range arrOfScheduleInt {
				scheduleMapped.Mon = append(scheduleMapped.Mon, arrOfScheduleInt[j])
			}
		case "Tue":
			scheduleSplitted := strings.Split(indexSplitter[1], "-") // 10:00-20:00
			arrOfScheduleInt := strScheduleToInt(scheduleSplitted)

			for j := range arrOfScheduleInt {
				scheduleMapped.Tue = append(scheduleMapped.Tue, arrOfScheduleInt[j])
			}
		case "Wed":
			scheduleSplitted := strings.Split(indexSplitter[1], "-") // 10:00-20:00
			arrOfScheduleInt := strScheduleToInt(scheduleSplitted)

			for j := range arrOfScheduleInt {
				scheduleMapped.Wed = append(scheduleMapped.Wed, arrOfScheduleInt[j])
			}
		case "Thu":
			scheduleSplitted := strings.Split(indexSplitter[1], "-") // 10:00-20:00
			arrOfScheduleInt := strScheduleToInt(scheduleSplitted)

			for j := range arrOfScheduleInt {
				scheduleMapped.Thu = append(scheduleMapped.Thu, arrOfScheduleInt[j])
			}
		case "Fri":
			scheduleSplitted := strings.Split(indexSplitter[1], "-") // 10:00-20:00
			arrOfScheduleInt := strScheduleToInt(scheduleSplitted)

			for j := range arrOfScheduleInt {
				scheduleMapped.Fri = append(scheduleMapped.Fri, arrOfScheduleInt[j])
			}
		case "Sat":
			scheduleSplitted := strings.Split(indexSplitter[1], "-") // 10:00-20:00
			arrOfScheduleInt := strScheduleToInt(scheduleSplitted)

			for j := range arrOfScheduleInt {
				scheduleMapped.Sat = append(scheduleMapped.Sat, arrOfScheduleInt[j])
			}
		}
	}

	return scheduleMapped
}

func strScheduleToInt(input []string) []int {
	result := []int{}

	for i := range input {
		strInputSplitted := strings.Split(input[i], ":")
		strToInt := 0
		var err error
		for j := range strInputSplitted {
			if len(strInputSplitted) > 1 {
				strToInt, err = strconv.Atoi(strInputSplitted[j] + strInputSplitted[j+1])
				if err != nil {
					log.Println("Error 97")
					log.Println(err)
				}
				break
			}
		}
		result = append(result, strToInt)
	}
	return result
}

func sortedScheduleOneDay(input []int) []int {
	sort.Slice(input, func(i, j int) bool {
		if input[i] < input[j] {
			return true
		}
		return false
	})

	return input
}

func weekCompleteSchedule(input []int) []int {
	result := []int{}

	for i := 0; i < 2500; i += 100 {
		result = append(result, i)
	}
	return result
}

func removeElement(input []int) []int {
	result := []int{}
	return result
}

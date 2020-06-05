package main

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
	"time"
)

func main() {
	input := "Sun 10:00-20:00\nFri 05:00-10:00\nFri 16:30-23:50\nSat 10:00-24:00\nSun 01:00-04:00\nSat 02:00-06:00\nTue 03:30-18:15\nTue 19:00-20:00\nWed 04:25-15:14\nWed 15:14-22:40\nThu 00:00-23:59\nMon 05:00-13:00\nMon 15:00-21:00"
	// "Sun 10:00-20:00 Fri 05:00-10:00 Fri 16:30-23:50 Sat 10:00-24:00 Sun 01:00-04:00 Sat 02:00-06:00 Tue 03:30-18:15 Tue 19:00-20:00 Wed 04:25-15:14 Wed 15:14-22:40 Thu 00:00-23:59 Mon 05:00-13:00 Mon 15:00-21:00" // 505 Tue 20:00 - Wed 04:25
	// "Mon 01:00-23:00 Tue 01:00-23:00 Wed 01:00-23:00 Thu 01:00-23:00 Fri 01:00-23:00 Sat 01:00-23:00 Sun 01:00-21:00" // 180 Sun 21:00 - 24:00
	output := solution(input)
	fmt.Println(output)
}

type mapped struct {
	Sun []string
	Mon []string
	Tue []string
	Wed []string
	Thu []string
	Fri []string
	Sat []string
}

type timeMapped struct {
	Sun []time.Time
	Mon []time.Time
	Tue []time.Time
	Wed []time.Time
	Thu []time.Time
	Fri []time.Time
	Sat []time.Time
}

func solution(str string) int {
	result := 0

	// split by \n
	stringSplitted := strings.Split(str, "\n")
	// assign value based on day to struct
	scheduleMapped := scheduleMapByDays(stringSplitted)
	valueScheduleMapped := reflect.ValueOf(scheduleMapped)
	values := make([]interface{}, valueScheduleMapped.NumField())
	for i := range values {
		values[i] = valueScheduleMapped.Field(i).Interface()
	}

	return result
}

func scheduleMapByDays(input []string) mapped {
	var scheduleMapped mapped

	for i := range input {
		indexSplitter := strings.Split(input[i], " ")
		tempArr := []string{}
		tempArrOfScheduleTime := []time.Time{}

		switch indexSplitter[0] {
		case "Sun":
			tempArrOfScheduleTime = timeFormatted(indexSplitter[1])
			tempArr = append(tempArr, indexSplitter[1])
			scheduleMapped.Sun = append(scheduleMapped.Sun, indexSplitter[1])
		case "Mon":
			scheduleMapped.Mon = append(scheduleMapped.Mon, indexSplitter[1])
		case "Tue":
			scheduleMapped.Tue = append(scheduleMapped.Tue, indexSplitter[1])
		case "Wed":
			scheduleMapped.Wed = append(scheduleMapped.Wed, indexSplitter[1])
		case "Thu":
			scheduleMapped.Thu = append(scheduleMapped.Thu, indexSplitter[1])
		case "Fri":
			scheduleMapped.Fri = append(scheduleMapped.Fri, indexSplitter[1])
		case "Sat":
			scheduleMapped.Sat = append(scheduleMapped.Sat, indexSplitter[1])
		}
		sort.Slice(tempArrOfScheduleTime, func(i, j int) bool {
			if tempArrOfScheduleTime[i].Unix() < tempArrOfScheduleTime[j].Unix() {
				return true
			}
			return false
		})
	}
	return scheduleMapped
}

func timeFormatted(input string) []time.Time {
	splittedSchedule := strings.Split(input, "-")
	strBeginningOfTheWeak := "2012-11-01T00:00:00+00:00"
	tempArrOfScheduleTime := []time.Time{}

	for j := range splittedSchedule {
		tempTime := splittedSchedule[j] + ":00"                                 // 10:00 + :00
		scheduleDaySplitted := strings.Split(strBeginningOfTheWeak, "T")        // 2012-11-01, 00:00:00+00:00
		exactTimeScheduleSplitted := strings.Split(scheduleDaySplitted[1], "+") // 00:00:00, 00:00
		scheduleDayInserted := fmt.Sprintf(scheduleDaySplitted[0] + "T" + tempTime + "+" + exactTimeScheduleSplitted[1])
		timeScheduleDayInserted, err := time.Parse(time.RFC3339, scheduleDayInserted)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(timeScheduleDayInserted)
		tempArrOfScheduleTime = append(tempArrOfScheduleTime, timeScheduleDayInserted)
	}
	return tempArrOfScheduleTime
}

// func (t *timeMapped) sortedOfScheduleByDay(input []time.Time) {
// 	for i := range input {
// 		*t = append(*t, input[i])
// 	}
// }

// iterate through struct, split

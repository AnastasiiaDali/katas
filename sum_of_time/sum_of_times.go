package main

import (
	"fmt"
	"strings"
	"time"
)

//Input: "12:32 34:01 15:23 9:27 55:22 25:56"
//Output: "2:32:41"

func main() {
	times := "12:32 34:01 15:23 9:27 55:22 25:56"
	fmt.Printf("Success %v", sumOfTime(times))
}

func sumOfTime(times string) string {
	var allSeconds float64
	array := strings.Split(times, " ")

	for _, time1 := range array {
		arrayOfTimes := strings.Split(time1, ":")
		duration, _ := time.ParseDuration(arrayOfTimes[0] + "m" + arrayOfTimes[1] + "s")
		allSeconds += duration.Seconds()
	}

	hours := int(allSeconds) / 3600
	minutes := int(allSeconds) / 60 % 60
	seconds := int(allSeconds) % 60

	result := fmt.Sprintf("%d:%d:%d", hours, minutes, seconds)

	return result
}

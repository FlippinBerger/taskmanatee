package main

import (
	"fmt"
	"time"
)

// DaysToSunday is a map denoting the number of days until the next Sunday. Allows us
// to get to the beginning of the week.
var DaysToSunday = map[string]int{
	"Sunday":    0,
	"Monday":    6,
	"Tuesday":   5,
	"Wednesday": 4,
	"Thursday":  3,
	"Friday":    2,
	"Saturday":  1,
}

// GetFileNameForToday will return a formatted string based on today's date
func GetFileNameForToday() string {
	today := time.Now()

	month := int(today.Month())
	day := today.Day()
	year := today.Year()

	return fmt.Sprintf("%d-%d-%d-tasks.txt", month, day, year)
}

// GetDateRangeForTime will return the time.Time values for Sunday and Saturday
// of the week containing day
func GetDateRangeForTime(day time.Time) (start time.Time, end time.Time) {
	start = day

	// Sunday = 0
	if start.Day() != 0 {
		day := time.Duration(start.Weekday())

		start = start.Add(time.Hour * 24 * day * -1)
	}

	end = start.Add(time.Hour * 24 * 6)

	return start, end
}

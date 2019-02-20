package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGetFileNameForToday(t *testing.T) {
	today := time.Now()

	month := int(today.Month())
	day := today.Day()
	year := today.Year()

	fileName := GetFileNameForToday()

	expected := fmt.Sprintf("%d-%d-%d-tasks.json", month, day, year)

	if fileName != expected {
		t.Errorf("Testing get file name for today.\nExpected: %s\nGot: %s", expected, fileName)
	}
}

func TestGetDateRangeForTime(t *testing.T) {
	// Sunday, February 4 2018
	// 2/4/2018 - 2/10/2018
	startOfWeek := time.Date(2018, 2, 4, 12, 0, 0, 0, time.UTC)
	start, end := GetDateRangeForTime(startOfWeek)

	expectedEnd := time.Date(2018, 2, 10, 12, 0, 0, 0, time.UTC)

	if !startOfWeek.Equal(start) {
		t.Errorf("Testing getDateRangeForTime for a Sunday.\n Start should be %v\nGot %v", startOfWeek, start)
	}

	if !expectedEnd.Equal(end) {
		t.Errorf("Testing getDateRangeForTime for a Sunday.\n End should be %v\nGot %v", expectedEnd, end)
	}

	// Saturday, February 9 2019
	// 2/3/2019 - 2/9/2019
	endOfWeek := time.Date(2019, 2, 9, 12, 0, 0, 0, time.UTC)
	start, end = GetDateRangeForTime(endOfWeek)

	expectedStart := time.Date(2019, 2, 3, 12, 0, 0, 0, time.UTC)

	if !expectedStart.Equal(start) {
		t.Errorf("Testing getDateRangeForTime for a Saturday.\n Start should be %v\nGot %v", expectedStart, start)
	}

	if !endOfWeek.Equal(end) {
		t.Errorf("Testing getDateRangeForTime for a Saturday.\n End should be %v\nGot %v", endOfWeek, end)
	}

	// Thursday, December 21 2017
	// 12/17/2017 - 12/23/2017
	midWeek := time.Date(2017, 12, 21, 12, 0, 0, 0, time.UTC)
	start, end = GetDateRangeForTime(midWeek)

	expectedStart = time.Date(2017, 12, 17, 12, 0, 0, 0, time.UTC)
	expectedEnd = time.Date(2017, 12, 23, 12, 0, 0, 0, time.UTC)

	if !expectedStart.Equal(start) {
		t.Errorf("Testing getDateRangeForTime for a Thursday.\n Start should be %v\nGot %v", expectedStart, start)
	}

	if !expectedEnd.Equal(end) {
		t.Errorf("Testing getDateRangeForTime for a Thursday.\n End should be %v\nGot %v", expectedEnd, end)
	}
}

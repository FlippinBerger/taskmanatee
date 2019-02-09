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

	expected := fmt.Sprintf("%d-%d-%d-tasks.txt", month, day, year)

	if fileName != expected {
		t.Errorf("Testing get file name for today.\nExpected: %s\nGot: %s", expected, fileName)
	}
}

func TestGetDateRangeForWeek(t *testing.T) {
	t.Skip("skipping til I figure out how to mock time")
	start, end = GetDateRangeForWeek()
}

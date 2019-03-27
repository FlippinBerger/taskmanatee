package main

import (
	"fmt"
)

// FormatTasks will output a CLI friendly UI for the given list of tasks
func FormatTasks(tasks *[]Task) {
	fmt.Println("Task List:")

	for i, task := range *tasks {
		fmt.Printf("%d. %v", i+1, &task)
	}
}

// OutputTasks gets the tasks for today, and uses OutputTasks to show them
func OutputTasks(tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println("You have no tasks recorded for today.")
	} else {
		FormatTasks(tasks)
	}
}

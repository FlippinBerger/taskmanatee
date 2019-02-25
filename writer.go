package main

import (
	"fmt"
)

// OutputTasks will output a CLI friendly UI for the given list of tasks
func OutputTasks(tasks *[]Task) {
	fmt.Println("Task List:")

	for i, task := range *tasks {
		fmt.Printf("%d. %s\n", i+1, task.Task)
	}
}

// OutputTasksForToday gets the tasks for today, and uses OutputTasks to show them
func OutputTasksForToday() {
	tasks, err := ReadFile(GetFileNameForToday())

	if err != nil || len(*tasks) == 0 {
		fmt.Println("You have no tasks recorded for today.")
	} else {
		OutputTasks(tasks)
	}
}

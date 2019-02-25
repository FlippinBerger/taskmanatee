package main

import (
	//"encoding/json"
	"fmt"
	"time"
)

// Task houses a complete task
type Task struct {
	Task          string     `json:"task"`
	Notes         []string   `json:"notes"`
	CreationTime  time.Time  `json:"creationTime"`
	Completed     bool       `json:"completed"`
	CompletedTime *time.Time `json:"completedTime"`
}

func newTask(todo string) *Task {
	task := Task{todo, nil, time.Now(), false, nil}
	return &task
}

// AddTasks will collect the tasks for the current day, and add the new tasks
// to it
func AddTasks(todos []string) (*[]Task, error) {
	tasks, err := ReadFile()

	if err != nil {
		taskErr := &TMError{fmt.Sprintf("There was an error reading the file: %v", err)}
		return nil, taskErr
	}

	// create a new task struct for each todo and add it to the slice of tasks
	taskSlice := *tasks

	for _, todo := range todos {
		task := newTask(todo)
		taskSlice = append(taskSlice, *task)
	}

	return &taskSlice, nil
}

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

// Tasks houses the entire Task list for a date
type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func newTask(todo string) *Task {
	task := Task{todo, nil, time.Now(), false, nil}
	return &task
}

// AddTasks will collect the tasks for the current day, and add the new tasks
// to it
func AddTasks(todos []string) (*[]Task, error) {
	fileName := GetFileNameForToday()

	fmt.Printf("file name is %s\n", fileName)
	fmt.Printf("todos are %v\n", todos)

	tasks, err := ReadFile(fileName)

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

	fmt.Printf("task Slice holds: %v", taskSlice)

	return &taskSlice, nil
}

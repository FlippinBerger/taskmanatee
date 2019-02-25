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

// Default string to be printed for a task
func (t Task) String() string {
	//TODO Add some sort of notion of completion here at the beginning
	info := "["

	if t.Completed {
		info += "X] "
	} else {
		info += " ] "
	}

	info += t.Task + "\n"

	if len(t.Notes) > 0 {
		for i, note := range t.Notes {
			info += fmt.Sprintf("\t%d. %s\n", i+1, note)
		}
	}

	return info
}

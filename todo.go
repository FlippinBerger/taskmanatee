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

//NewTask is a construtor for a new task.
func NewTask(todo string, notes []string) *Task {
	task := Task{todo, notes, time.Now(), false, nil}
	return &task
}

// for now we're just adding tasks for today.
func addTasks(todos []string) []Task {
	fileName := GetFileNameForToday()

	fmt.Printf("file name is %s\n", fileName)

	// TODO just ignoring the error here for now, I'll come back and deal with
	// it later
	tasksStruct, err := ReadFile(fileName)

	if err != nil {
		fmt.Printf("There was an error reading the file: %v\n", err)
	}

	fmt.Println("got the tasks struct")
	fmt.Printf("struct is %v\n", tasksStruct)
	fmt.Println("made it here after struct")

	tasks := tasksStruct.Tasks

	fmt.Println("got the tasks")

	// create a new task struct for each todo and add it to the slice of tasks
	for _, todo := range todos {
		task := Task{Task: todo}
		tasks = append(tasks, task)
	}

	return tasks
}

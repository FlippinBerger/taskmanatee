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

// AppendTasks will append 2 slices of tasks togeter
func AppendTasks(tasks *[]Task, newTasks *[]Task) *[]Task {
	newSlice := append(*tasks, *newTasks...)
	return &newSlice
}

// CreateTask will take do all the work of adding a string to a list of tasks
func CreateTask(tasks *[]Task, todo string) *[]Task {
	task := newTask(todo)
	return AppendTasks(tasks, &[]Task{*task})
}

// DeleteTask will remove the task at index
func DeleteTask(tasks *[]Task, index int) {
	if len(*tasks) == 1 {
		*tasks = nil
	} else if index == len(*tasks)-1 {
		*tasks = (*tasks)[:index]
	} else {
		*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
	}
}

// Complete will mark a task as complete
func (t *Task) Complete() {
	t.Completed = true
}

// Uncomplete will mark a task uncomplete
func (t *Task) Uncomplete() {
	t.Completed = false
}

// AddNote adds the provided note to the task
func (t *Task) AddNote(note string) {
	t.Notes = append(t.Notes, note)
}

// DeleteNote will remove the note at index
func (t *Task) DeleteNote(index int) {
	if len(t.Notes) == 1 {
		t.Notes = nil
	} else if index == len(t.Notes)-1 {
		t.Notes = t.Notes[:index]
	} else {
		t.Notes = append(t.Notes[:index], t.Notes[index+1:]...)
	}
}

// Default string to be printed for a task
func (t *Task) String() string {
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

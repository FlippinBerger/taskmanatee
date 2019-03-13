package main

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	todo := "Do the thing."
	task := newTask(todo)

	if task.Task != todo {
		t.Errorf("task intialized with incorrect todo")
	}

	if task.Completed {
		t.Errorf("task shouldn't be completed on initialization")
	}
}

func TestAppendTasksBasic(t *testing.T) {
	t1 := newTask("Existing task 1")
	t2 := newTask("Existing task 2")

	existingTasks := &[]Task{*t1, *t2}

	t3 := newTask("task to append")

	existingTasks = AppendTasks(existingTasks, &[]Task{*t3})

	if len(*existingTasks) != 3 {
		t.Errorf("the task array should contain 3 tasks after appending, but it contains %d", len(*existingTasks))
	}

	finalTask := (*existingTasks)[2]

	if finalTask.Task != "task to append" {
		t.Errorf("The last task should be the task just appended")
	}
}

func TestDeleteTask(t *testing.T) {
	t1 := newTask("Existing task 1")
	t2 := newTask("Existing task 2")
	t3 := newTask("Existing task 3")

	existingTasks := []Task{*t1, *t2, *t3}

	//delete the middle task
	DeleteTask(&existingTasks, 1)

	if len(existingTasks) != 2 {
		t.Errorf("The count after deletion should be 2, but it's %d", len(existingTasks))
	}

	t1 = &existingTasks[0]
	if t1.Task != "Existing task 1" {
		t.Errorf("first task is incorrect")
	}

	t2 = &existingTasks[1]
	if t2.Task != "Existing task 3" {
		t.Errorf("second task is incorrect")
	}

	DeleteTask(&existingTasks, 1)
	if len(existingTasks) != 1 {
		t.Errorf("the count should be 1, but it's %d", len(existingTasks))
	}

	t1 = &existingTasks[0]

	if t1.Task != "Existing task 1" {
		t.Errorf("first task is incorrect")
	}
}

func TestCompleteTask(t *testing.T) {
	t1 := newTask("Existing task 1")

	if t1.Completed {
		t.Errorf("task shouldn't be completed on initialization")
	}

	t1.Complete()

	if !t1.Completed {
		t.Errorf("task should be complete after calling complete")
	}

	t1.Uncomplete()
	if t1.Completed {
		t.Errorf("task shouldn't be completed after calling uncomplete")
	}
}

func TestAddNote(t *testing.T) {
	t1 := newTask("task")

	if len(t1.Notes) != 0 {
		t.Errorf("task shouldn't have notes before adding them")
	}

	t1.AddNote("do it")

	if len(t1.Notes) != 1 {
		t.Errorf("length of notes should be 1, but it's %d", len(t1.Notes))
	}

	t1.AddNote("yeehaw")

	if len(t1.Notes) != 2 {
		t.Errorf("length of notes should be 2, but it's %d", len(t1.Notes))
	}

	if t1.Notes[0] != "do it" {
		t.Errorf("The notes are out of order after addition")
	}

	t1.AddNote("Note 3")

	t1.DeleteNote(1)

	if len(t1.Notes) != 2 {
		t.Errorf("after deletion length of notes should be 2, but it's %d", len(t1.Notes))
	}

	if t1.Notes[1] != "Note 3" {
		t.Errorf("note 1 should be \"Note 3\", but instead it's %s", t1.Notes[1])
	}
}

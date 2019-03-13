package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

var path = ""

func getFilePath() string {
	if path != "" {
		return path
	}

	user, _ := user.Current()
	dir := user.HomeDir

	path = dir + "/.taskManTasks.json"

	return path
}

func createFile() {
	if !FileExists() {
		os.Create(getFilePath())
	} else {
		fmt.Println("Unable to create file because it already exists")
	}
}

// FileExists returns true if the file at path exists
func FileExists() bool {
	_, err := os.Stat(getFilePath())

	return !os.IsNotExist(err)
}

func deleteFile() {
	//TODO pop the confirmation to make sure they want to delete the file
	os.Remove(getFilePath())
}

// ReadFile will try to open and parse out the given file
func ReadFile() (*[]Task, error) {
	if !FileExists() {
		fmt.Println("creating the file")
		createFile()
	}
	filePath := getFilePath()
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		// handle error here
		tmError := &TMError{fmt.Sprintf("Error reading from %s:, %v", filePath, err)}
		return nil, tmError
	}

	var taskSlice []Task
	tasks := &taskSlice

	// if the file is empty, just return empty tasks
	if len(data) == 0 {
		return tasks, nil
	}

	// now we need to parse out the data from the file
	err = json.Unmarshal([]byte(data), tasks)

	if err != nil {
		tmError := &TMError{fmt.Sprintf("Unable to parse file due to error: %v", err)}
		return nil, tmError
	}

	return tasks, nil
}

// OverwriteFile will format a Tasks struct to json to be written to the task
// file
func OverwriteFile(tasks *[]Task) error {
	data, err := json.Marshal(&tasks)

	if err != nil {
		return err
	}

	ioutil.WriteFile(getFilePath(), data, 0666)
	return nil
}

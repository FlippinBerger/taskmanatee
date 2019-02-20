package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// File is an object to denote a File of tasks
type File struct {
	Name  string
	Tasks []Task
}

func getFileNameFromPath(path string) string {
	slice := strings.Split(path, "/")
	return slice[len(slice)-1]
}

func createFileAtPath(path string) {
	filename := getFileNameFromPath(path)

	if !FileExists(path) {
		os.Create(filename)
	} else {
		fmt.Println("Unable to create file because it already exists")
	}
}

// FileExists returns true if the file at path exists
func FileExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func deleteFile(file string) {
	//TODO pop the confirmation to make sure they want to delete the file
	os.Remove(file)
}

// ReadFile will try to open and parse out the given file
func ReadFile(fileName string) (*[]Task, error) {
	dir, _ := os.Getwd()
	fullPath := dir + "/" + fileName

	if !FileExists(fullPath) {
		createFileAtPath(fullPath)
	}
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		// handle error here
		tmError := &TMError{fmt.Sprintf("Error reading from %s:, %v", fileName, err)}
		return nil, tmError
	}

	taskSlice := make([]Task, 5, 5)
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
	fmt.Println("Overwriting file here.")
	fileName := GetFileNameForToday()

	data, err := json.Marshal(&tasks)

	if err != nil {
		return err
	}

	ioutil.WriteFile(fileName, data, 0666)
	return nil
}

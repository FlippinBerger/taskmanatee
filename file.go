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
func ReadFile(fileName string) (*Tasks, error) {
	dir, _ := os.Getwd()
	fullPath := dir + "/" + fileName

	if !FileExists(fullPath) {
		createFileAtPath(fullPath)
	}
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		// handle error here
		return nil, err
	}

	tasks := Tasks{}

	err = json.Unmarshal([]byte(file), &tasks)

	if err != nil {
		//handle error again
		return nil, err
	}

	return &tasks, nil
}

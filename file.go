package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// File is an object to denote a File of tasks
type File struct {
	Name  string
	Tasks []Task
}

func createFile(dir string, force bool) {
	filename := GetFileNameForToday()
	fmt.Println(filename)

	// check if the file already exists
	path := dir + "/" + filename

	if !FileExists(path) {
		os.Create(filename)
	} else if force {
		//TODO make sure the user is sure they want to overwrite the file
		//for now just overwrite it with a slightly different name
		//message := fmt.Sprintf("Are you sure you want to overwrite %s?", filename)

		filename = "forced-" + filename
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

func getFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	return file, nil
}

// ReadFile will try to open and parse out the given file
func ReadFile(fileName string) (*Tasks, error) {
	dir, _ := os.Getwd()
	fullPath := dir + "/" + fileName

	if !FileExists(fullPath) {
		createFile(dir, false)
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

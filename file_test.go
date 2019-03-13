package main

import (
	"os"
	"os/user"
	"testing"
)

func getPath() string {
	if path != "" {
		return path
	}

	user, _ := user.Current()
	dir := user.HomeDir

	path = dir + "/.taskManTasks.json"

	return path
}

func TestCreateFile(t *testing.T) {
	if FileExists() {
		t.Errorf("Testing createFile, and the file doesn't exist")
	}

	createFile()

	if !FileExists() {
		t.Errorf("Testing createFile, and the file wasn't created")
	} else {
		// clean up after ourselves, this test was successful
		os.Remove(getPath())
	}
}

func TestFileExists(t *testing.T) {
	path := getPath()

	exists := FileExists()

	if exists {
		t.Errorf("Testing FileExists, and if returned true for a file not yet created.")
	}

	os.Create(path)

	exists = FileExists()

	if !exists {
		t.Errorf("Testing FileExists, and if returned false for a file just created.")
	} else {
		//clean up
		os.Remove(path)
	}
}

func TestDeleteFile(t *testing.T) {
	path := getPath()
	os.Create(path)

	deleteFile()

	exists := FileExists()

	if exists {
		t.Errorf("Testing DeleteFile, and it didn't delete the file.")
	}
}

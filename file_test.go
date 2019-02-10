package main

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	// get the pwd
	dir, _ := os.Getwd()

	// get the file name for today
	fileName := "test-" + GetFileNameForToday()

	// check if the file already exists
	path := dir + "/" + fileName

	if FileExists(path) {
		t.Errorf("Testing createFile, and the file for today already exists")
	}

	createFileAtPath(path)

	if !FileExists(path) {
		t.Errorf("Testing createFile, and the file wasn't created")
	} else {
		// clean up after ourselves, this test was successful
		os.Remove(fileName)
	}
}

func TestFileExists(t *testing.T) {
	dir, _ := os.Getwd()

	fileName := "fileExistsTest.txt"

	path := dir + "/" + fileName

	exists := FileExists(path)

	if exists {
		t.Errorf("Testing FileExists, and if returned true for a file not yet created.")
	}

	os.Create(fileName)

	exists = FileExists(path)

	if !exists {
		t.Errorf("Testing FileExists, and if returned false for a file just created.")
	} else {
		//clean up
		os.Remove(fileName)
	}
}

func TestDeleteFile(t *testing.T) {
	dir, _ := os.Getwd()
	fileName := "deleteFileTest.txt"

	path := dir + "/" + fileName

	os.Create(fileName)

	deleteFile(fileName)

	exists := FileExists(path)

	if exists {
		t.Errorf("Testing DeleteFile, and it didn't delete the file.")
	}
}

func TestGetFileNameFromPath(t *testing.T) {
	path := "home/dir/file.txt"

	fileName := getFileNameFromPath(path)

	expected := "file.txt"

	if fileName != expected {
		t.Errorf("Testing getFileNameFromPath for path %s.\nExpected %s\nGot %s\n", path, expected, fileName)
	}
}

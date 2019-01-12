package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// knowing where we are in the filesystem is vital for this tool. If we
	// can't get the PWD, we need to get out.
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("we have some serious issues and can't get the current working directory, exiting")
		return
	}

	// Parse the command line flags.
	create := flag.Bool("create", false, "Pass create to create a new task file.")
	force := flag.Bool("f", false, "Pass f to overwrite files that already exist.")
	delete := flag.String("d", "", "Pass d to delete a specified task file.")
	//num := flag.Int("n", 0, "Pass n to use a day different from the current day.")
	//dir := flag.String("dir", currentDir, "Pass a string to dir to specifiy where to perform the action.")
	//week := flag.Bool("week", false, "Pass week to show the tasks for the entire week Sunday-Saturday")

	flag.Parse()

	if *create {
		fmt.Println("Hell yeah, we in there.")
		createFile(currentDir, *force)
	}

	if *delete != "" {
		deleteFile(*delete)
	}

	// if no arguments are provided,
	if len(os.Args) == 1 {
		fmt.Println("Showing your tasks for today.")
		showTasks()
	}

	//fmt.Printf("dir is %s\n", *dir)
}

func getFileNameForToday() string {
	today := time.Now()
	fmt.Println(today)

	month := int(today.Month())
	fmt.Println(month)

	day := today.Day()
	fmt.Println(day)

	year := today.Year()
	fmt.Println(year)

	return fmt.Sprintf("%d-%d-%d-tasks.txt", month, day, year)
}

func createFile(dir string, force bool) {
	filename := getFileNameForToday()
	fmt.Println(filename)

	// check if the file already exists
	path := dir + "/" + filename

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(filename)
	} else if force {
		//TODO make sure the user is sure they want to overwrite the file
		//for now just overwrite it with a slightly different name
		filename = "forced-" + filename
		os.Create(filename)
	} else {
		fmt.Println("Unable to create file because it already exists")
	}
}

func deleteFile(file string) {
	//TODO pop the confirmation to make sure they want to delete the file
	os.Remove(file)
}

func showTasks() {
	fmt.Println("Show tasks has been called.")
}

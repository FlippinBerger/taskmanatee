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
	task := flag.Bool("t", false, "Pass -t at the end to pass in a list of tasks as tailing args.")

	flag.Parse()

	if *create {
		createFile(currentDir, *force)
	}

	if *delete != "" {
		deleteFile(*delete)
	}

	// if no arguments are provided,
	if len(os.Args) == 1 {
		showTodaysTasks()
	}

	if *task {
		fmt.Println("tail: ", flag.Args())
		fmt.Println("count: ", len(flag.Args()))
	}
}

func getFileNameForToday() string {
	today := time.Now()

	month := int(today.Month())
	day := today.Day()
	year := today.Year()

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
		//message := fmt.Sprintf("Are you sure you want to overwrite %s?", filename)

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

func showTodaysTasks() {
	fmt.Println("Showing all your tasks for the day.")
}

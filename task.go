package main

import (
	"flag"
	"fmt"
	"os"
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
	task := flag.Bool("t", false, "Pass -t at the end to pass in a list of tasks as trailing args.")
	//week := flag.Bool("w", false, "Pass week to show the tasks for the entire week Sunday-Saturday")

	delete := flag.String("d", "", "Pass d to delete a specified task file.")
	// TODO this is going to need to be handled differently
	//num := flag.Int("n", 0, "Pass n to use a day different from the current day.")
	//dir := flag.String("dir", currentDir, "Pass a string to dir to specifiy where to perform the action.")

	flag.Parse()

	// if no arguments are provided, show the tasks for today.
	if len(os.Args) == 1 {
		showTodaysTasks()
	}

	if *create {
		fullPath := currentDir + GetFileNameForToday()
		createFileAtPath(fullPath)
	}

	if *delete != "" {
		deleteFile(*delete)
	}

	if *task {
		fmt.Println("tail: ", flag.Args())
		fmt.Println("count: ", len(flag.Args()))
		tasks, err := AddTasks(flag.Args())

		fmt.Printf("tasks are: %v\n", tasks)

		if err != nil {
			fmt.Println(err)
			fmt.Println("we definitely hit an error")
			return
		}

		// Commit the tasks to the file
		OverwriteFile(tasks)
	}
}

func showTodaysTasks() {
	fmt.Println("Showing all your tasks for the day.")
}

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse the command line flags.
	task := flag.Bool("t", false, "Pass -t at the end to pass in a list of tasks as trailing args.")
	//week := flag.Bool("w", false, "Pass week to show the tasks for the entire week Sunday-Saturday")

	delete := flag.String("d", "", "Pass d to delete a specified task file.")
	// TODO this is going to need to be handled differently
	//num := flag.Int("n", 0, "Pass n to use a day different from the current day.")
	//dir := flag.String("dir", currentDir, "Pass a string to dir to specifiy where to perform the action.")

	flag.Parse()

	// if no arguments are provided, show the tasks for today.
	if len(os.Args) == 1 {
		OutputTasksForToday()
	}

	if *delete != "" {
		deleteFile(*delete)
	}

	if *task {
		tasks, err := AddTasks(flag.Args())

		if err != nil {
			fmt.Println(err)
			fmt.Println("we definitely hit an error")
			return
		}

		OutputTasks(tasks)

		// Commit the tasks to the file
		OverwriteFile(tasks)
	}
}

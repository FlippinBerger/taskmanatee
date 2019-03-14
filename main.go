package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse the command line flags.
	task := flag.Bool("t", false, "Pass -t at the end to pass in a list of tasks as trailing args.")

	delete := flag.Bool("d", false, "Pass d to delete the stored task file.")

	flag.Parse()

	// if no arguments are provided, show the tasks for today.
	if len(os.Args) == 1 {
		OutputTasks()
	}

	if *delete {
		deleteFile()
	}

	if *task {
		tasks, err := AddTasks(flag.Args())

		if err != nil {
			fmt.Println(err)
			fmt.Println("we definitely hit an error")
			return
		}

		FormatTasks(tasks)

		// Commit the tasks to the file
		OverwriteFile(tasks)
	}
}

package main

import (
	"bufio"
	//	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	// get the existing tasks if they're saved on disk
	var tasks *[]Task

	tasks, err := ReadFile()

	// if the file wasn't able to be read, allocate some memory
	if err != nil {
		var taskSlice []Task
		tasks = &taskSlice
	}

	// capture Ctrl+C to write the info back to file
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// kick off thread to wait for interrupts
	// this blocking can be removed once we're out of the file system for
	// storage
	go func() {
		_ = <-sigs
		done <- true
	}()

	scanner := bufio.NewScanner(os.Stdin)

	// infinitely loop and wait for user input
	for {
		// show the tasks
		OutputTasks(tasks)

		// save and exit if we've gotten our interrupt signal
		select {
		case _ = <-done:
			fmt.Println("Overwriting and killing the thread")
			OverwriteFile(tasks)
			os.Exit(1)
		default:
			// Grab the user's input, and pass it to the handler
			scanner.Scan()
		}

		tasks = handleInput(scanner.Text(), tasks)
	}
}

func handleInput(text string, tasks *[]Task) *[]Task {
	// strip all leading and trailing whitespace
	text = strings.TrimSpace(text)

	// just pressing enter without input will write the file and end the
	// program
	if len(text) == 0 {
		OverwriteFile(tasks)
		os.Exit(1)
	}

	words := strings.Fields(text)

	if len(words) < 2 {
		// printUsage here too
		return tasks
	}

	command := words[0]

	// set this flag if any formatting was incorrect through parsing
	printUsage := false

	switch command {
	case "c":
		fmt.Println("creating")
		entity := words[1]

		switch entity {
		case "t":
			task := text[4:]
			fmt.Println(task)

			tasks = CreateTask(tasks, task)
		case "n":
			//create note
			if i, err := strconv.Atoi(words[2]); err == nil {
				if i < 1 || i > len(*tasks) {
					printUsage = true
				} else {
					// create the note
					(*tasks)[i-1].AddNote(text[6:])
				}
			} else {
				printUsage = true
			}
		default:
			printUsage = true
		}
	case "d":
		fmt.Println("deleting")
		entity := words[1]

		switch entity {
		case "t":
			if i, err := strconv.Atoi(words[2]); err == nil {
				if i < 1 || i > len(*tasks) {
					printUsage = true
				} else {
					// delete the task
					DeleteTask(tasks, i-1)
				}
			} else {
				printUsage = true
			}
			//delete task
		case "n":
			//delete note
			if i, err := strconv.Atoi(words[2]); err == nil {
				if i < 1 || i > len(*tasks) {
					printUsage = true
				} else {
					if j, err := strconv.Atoi(words[3]); err == nil {
						if j < 1 || j > len(task.Notes) {
							printUsage = true
						} else {
							(*tasks)[i-1].DeleteNote(j - 1)
						}
					}
				}
			} else {
				printUsage = true
			}
		default:
			printUsage = true
		}
	case "x":
		fmt.Println("completing")
		if i, err := strconv.Atoi(words[1]); err == nil {
			if i < 1 || i > len(*tasks) {
				printUsage = true
			} else {
				// complete the task
				(*tasks)[i-1].Complete()
			}
		} else {
			printUsage = true
		}
	case "u":
		fmt.Println("uncompleting")
		if i, err := strconv.Atoi(words[1]); err == nil {
			if i < 1 || i > len(*tasks) {
				printUsage = true
			} else {
				// complete the task
				(*tasks)[i-1].Uncomplete()
			}
		} else {
			printUsage = true
		}
	case "s":
		//sort command
		fmt.Println("sorting")
	case "f":
		//filter command
		fmt.Println("filtering")
	default:
		//unknown command, return error message and usage
		fmt.Println("Unknown command")
		printUsage = true
	}

	if printUsage {
		// say formatting was incorrect and return the usage block
		fmt.Println("printing usage")
	}

	return tasks
}

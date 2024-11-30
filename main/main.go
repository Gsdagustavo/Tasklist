package main

import (
	"errors"
	"fmt"
	"os"
)

type Task struct {
	name        string
	isCompleted bool
}

// removes the first argument (element) of a slice
func removeFirstArg(args []string) ([]string, error) {
	if len(args) > 1 {
		return args[1:], nil // exclude the first argument
	}

	// returns an empty slice if no arguments are given
	return []string{}, errors.New("removeFirstArgs func expects at least 2 arguments, but less were given")
}

var commands = map[int]string{
	1: "add",
	2: "remove",
	3: "list",
	4: "change",
}

func validateCommand(args []string) (int, error) {
	if len(args) >= 1 {
		switch args[0] {
		case "add":
			return 1, nil
		case "remove":
			return 2, nil
		case "list":
			return 3, nil
		case "change":
			return 4, nil
		}
	}

	return -1, errors.New("validateCommands expect an array of strings, but an empty one was given")
}

// global slice to store tasks
var tasks []Task

const MAX_TASKS int = 10

// adds a task to the tasks slice
func addTask(args []string) {

	if len(args) >= MAX_TASKS {
		fmt.Println("The task list is full! Cannot add more tasks.")
		return
	}

	if len(args) < 2 {
		fmt.Println("addTask expects a task name as an argument.")
		return
	}

	// creates a new task and appends it to the slice
	task := Task{name: args[1], isCompleted: false}
	tasks = append(tasks, task)
	fmt.Printf("Task \"%v\" added to the tasklist.\n", task.name)
}

func listTasks() {
	if len(tasks) <= 0 {
		fmt.Println("The list is empty.")
		return
	}

	fmt.Println("Listing all tasks:")

	for i := 1; i <= len(tasks); i++ {
		fmt.Printf("%v. %v", i, tasks[i])
	}
}

func main() {

	args, err := removeFirstArg(os.Args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	commandId, _ := validateCommand(args)

	switch commandId {
	case 1:
		addTask(args)
	case 3:
		listTasks()
	}

}

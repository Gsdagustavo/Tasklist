package main

import (
	"errors"
	"fmt"
	"os"
)

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
	fmt.Printf("commands[0]: %v\n", commands[0])
	fmt.Printf("commands[1]: %v\n", commands[1])

	fmt.Printf("args: %v\n", args)

	if len(args) > 0 {
		for key, value := range commands {
			if args[0] == value {
				return key, nil
			}
		}
	}

	return -1, errors.New("at least one argument is expected, but none were given")
}

// global slice to store tasks
var tasks []Task

const MAX_TASKS int = 10

// adds a task to the tasks slice
func addTask(args []string) {

	// file, err := os.Open("data.JSON")

	// if err != nil {
	// 	err.Error("could not open the data.JSON file")
	// }

	test()

	if len(args) >= MAX_TASKS {
		fmt.Println("The task list is full! Cannot add more tasks.")
		return
	}

	if len(args) < 2 {
		fmt.Println("addTask expects a task name as an argument.")
		return
	}

	// creates a new task and appends it to the slice
	// task := Task{name: args[1], isCompleted: false}
	// tasks = append(tasks, task)
	// fmt.Printf("Task \"%v\" added to the tasklist.\n", task.name)
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

	fmt.Printf("commandId: %v\n", commandId)

	// switch commandId {
	// case 1:
	// 	addTask(args)
	// case 3:
	// 	listTasks()
	// }

}

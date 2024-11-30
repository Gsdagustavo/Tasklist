package main

import (
	"errors"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	importance  int
	id          uint16
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

func main() {

	args, err := removeFirstArg(os.Args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(args)

}

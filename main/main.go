package main

import (
	"os"
)

type Task struct {
	name        string
	description string
	importance  int
	id          uint16
	isCompleted bool
}

func main() {

	args := os.Args

	for i := range args {

	}

}

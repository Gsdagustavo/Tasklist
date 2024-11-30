package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Task represents a task in the task list
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func test() {
	// Example task list
	tasks := []Task{}

	// Serialize (marshal) tasks to JSON
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return
	}

	// Write JSON to a file
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Tasks successfully written to tasks.json")

	// Read JSON from the file and deserialize (unmarshal) it back into a slice of Task structs
	file, err = os.Open("tasks.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var loadedTasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&loadedTasks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Loaded tasks:", loadedTasks)
}

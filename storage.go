package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func loadTasks() []Task {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return []Task{}
	}
	return tasks
}

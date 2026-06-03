package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Title     string
	Completed bool
}

func addTask(tasks []Task) []Task {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Task : ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input", err)
		return tasks
	}
	title = strings.TrimSpace(title)
	newTask := Task{
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, newTask)

	fmt.Println("Task added successfully")

	return tasks
}

func viewTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No task available")
		return
	}

	for index, task := range tasks {
		if task.Completed {
			fmt.Printf("%d. [x] %s\n", index+1, task.Title)
		} else {
			fmt.Printf("%d. [ ] %s\n", index+1, task.Title)

		}
	}
}

func deleteTask(tasks []Task) []Task {
	var deleteNo int

	viewTasks(tasks)

	if len(tasks) == 0 {
		fmt.Println("No Tasks available")
		return tasks
	}

	fmt.Printf("Enter task number : ")
	fmt.Scan(&deleteNo)

	index := deleteNo - 1

	if deleteNo <= 0 || deleteNo > len(tasks) {
		fmt.Println("No Such Tasks")
	} else {
		tasks = append(tasks[:index], tasks[index+1:]...)
		fmt.Println("Task deleted successfully")
	}
	return tasks
}

func completeTask(tasks []Task) []Task {
	if len(tasks) == 0 {
		fmt.Println("No task available")
		return tasks
	}

	viewTasks(tasks)

	var taskNoC int

	fmt.Println("Enter task number to complete : ")
	fmt.Scan(&taskNoC)

	if taskNoC <= 0 || taskNoC > len(tasks) {
		fmt.Println("Invalid task number")
		return tasks
	}

	index := taskNoC - 1

	tasks[index].Completed = true

	fmt.Println("Task marked as completed")
	return tasks
}

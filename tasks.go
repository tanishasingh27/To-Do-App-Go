package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func readInt(reader *bufio.Reader) (int, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(line))
}

type Task struct {
	Title     string
	Completed bool
}

func createTask(title string) Task {
	return Task{
		Title:     title,
		Completed: false,
	}
}

func markTaskComplete(tasks []Task, index int) []Task {
	tasks[index].Completed = true
	return tasks
}

func addTask(reader *bufio.Reader, tasks []Task) []Task {
	fmt.Println("Enter Task : ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input", err)
		return tasks
	}
	title = strings.TrimSpace(title)
	if title == "" {
		fmt.Println("Task title cannot be empty")
		return tasks
	}
	newTask := createTask(title)

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

func deleteTask(reader *bufio.Reader, tasks []Task) []Task {
	viewTasks(tasks)

	if len(tasks) == 0 {
		fmt.Println("No Tasks available")
		return tasks
	}

	fmt.Printf("Enter task number : ")
	deleteNo, err := readInt(reader)
	if err != nil {
		fmt.Println("Invalid task number")
		return tasks
	}

	index := deleteNo - 1

	if deleteNo <= 0 || deleteNo > len(tasks) {
		fmt.Println("No Such Tasks")
	} else {
		tasks = append(tasks[:index], tasks[index+1:]...)
		fmt.Println("Task deleted successfully")
	}
	return tasks
}

func completeTask(reader *bufio.Reader, tasks []Task) []Task {
	if len(tasks) == 0 {
		fmt.Println("No task available")
		return tasks
	}

	viewTasks(tasks)

	fmt.Println("Enter task number to complete : ")
	taskNoC, err := readInt(reader)
	if err != nil {
		fmt.Println("Invalid task number")
		return tasks
	}

	if taskNoC <= 0 || taskNoC > len(tasks) {
		fmt.Println("Invalid task number")
		return tasks
	}

	index := taskNoC - 1

	tasks = markTaskComplete(tasks, index)

	fmt.Println("Task marked as completed")
	return tasks
}

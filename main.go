package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tasks := loadTasks()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--TO_DO LIST--")

	for {
		fmt.Printf("\n1.Add task\n2.View Task\n3.Complete Task\n4.Delete Task\n5.Exit\n")
		fmt.Printf("Enter your choice : ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		choice, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			tasks = addTask(reader, tasks)
			saveTasks(tasks)
		case 2:
			viewTasks(tasks)
		case 3:
			tasks = completeTask(reader, tasks)
			saveTasks(tasks)
		case 4:
			tasks = deleteTask(reader, tasks)
			saveTasks(tasks)
		case 5:
			fmt.Println("Exiting..")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}

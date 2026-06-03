package main

import "fmt"

func main() {
	var choice uint
	tasks := loadTasks()

	fmt.Println("--TO_DO LIST--")

	for {
		fmt.Printf("\n1.Add task\n2.View Task\n3.Complete Task\n4.Delete Task\n5.Exit\n")
		fmt.Printf("Enter your choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tasks = addTask(tasks)
			saveTasks(tasks)
		case 2:
			viewTasks(tasks)
		case 3:
			tasks = completeTask(tasks)
			saveTasks(tasks)
		case 4:
			tasks = deleteTask(tasks)
			saveTasks(tasks)
		case 5:
			fmt.Println("Exiting..")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}

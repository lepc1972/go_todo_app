package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var taskItems = []string{}

func main() {
	fmt.Println("##### Welcome to our todolist app #####")
	for {
		printMenu()
		choice := getUserChoice()
		switch choice {
		case "1":
			checkAndDeleteTasks()
		case "2":
			addTask()
		case "3":
			fmt.Println("Exiting the app. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Show tasks and check if executed")
	fmt.Println("2. Add a task")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice: ")
}

func getUserChoice() string {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func addTask() {
	fmt.Print("\nEnter the task you want to add: ")
	reader := bufio.NewReader(os.Stdin)
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
	if task != "" {
		taskItems = append(taskItems, task)
		fmt.Println("Task added successfully!")
	} else {
		fmt.Println("Task cannot be empty!")
	}
}

func checkAndDeleteTasks() {
	if len(taskItems) == 0 {
		fmt.Println("\nNo tasks yet. Add some!")
		return
	}

	fmt.Println("\nList of my todos:")
	remainingTasks := []string{} // Temporary slice for tasks that are not executed

	reader := bufio.NewReader(os.Stdin)
	for i, task := range taskItems {
		fmt.Printf("Task %d: %s\n", i+1, task)
		fmt.Print("Was this task executed? (yes/no): ")
		response, _ := reader.ReadString('\n')
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "yes" {
			fmt.Printf("Task '%s' marked as executed and deleted.\n", task)
		} else if response == "no" {
			remainingTasks = append(remainingTasks, task)
		} else {
			fmt.Println("Invalid input. Keeping the task in the list.")
			remainingTasks = append(remainingTasks, task)
		}
	}

	// Update the task list with remaining tasks
	taskItems = remainingTasks

	// Show updated tasks
	fmt.Println("\nUpdated list of tasks:")
	if len(taskItems) == 0 {
		fmt.Println("No remaining tasks!")
	} else {
		for i, task := range taskItems {
			fmt.Printf("Task %d: %s\n", i+1, task)
		}
	}
}

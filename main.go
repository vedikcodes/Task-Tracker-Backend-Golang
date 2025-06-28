package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var taskFile = "tasks.json"

// Function to load tasks from the file
func loadTasks() ([]Task, error) {
	// Check if the file exists and read it
	file, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Return empty list if the file doesn't exist
			return []Task{}, nil
		}
		return nil, err
	}

	// Handle the case where the file is empty or contains malformed data
	if len(file) == 0 {
		return []Task{}, nil
	}

	// Deserialize JSON data into tasks
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tasks.json: %w", err)
	}
	return tasks, nil
}

// Function to save tasks to the file
func saveTasks(tasks []Task) error {
	// Serialize tasks and write to file
	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, file, 0644)
}

// Add a new task
func addTask(description string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Generate a unique ID for the new task
	id := len(tasks) + 1
	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Append the new task and save
	tasks = append(tasks, task)
	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)\n", id)
}

// List tasks based on their status
func listTasks(status string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for _, task := range tasks {
		if status == "" || task.Status == status {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

// Update the task description
func updateTask(id int, newDescription string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			err = saveTasks(tasks)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
				return
			}
			fmt.Println("Task updated successfully")
			return
		}
	}
	fmt.Println("Task not found")
}

// Delete a task
func deleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			err = saveTasks(tasks)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
				return
			}
			fmt.Println("Task deleted successfully")
			return
		}
	}
	fmt.Println("Task not found")
}

// Mark a task as in-progress
func markInProgress(id int) {
	updateTaskStatus(id, "in-progress")
}

// Mark a task as done
func markDone(id int) {
	updateTaskStatus(id, "done")
}

// Helper function to update the task status
func updateTaskStatus(id int, status string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			err = saveTasks(tasks)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
				return
			}
			fmt.Printf("Task marked as %s\n", status)
			return
		}
	}
	fmt.Println("Task not found")
}

// Helper function to parse task ID from string to int
func parseID(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID:", idStr)
		os.Exit(1)
	}
	return id
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <task-description>")
			return
		}
		addTask(os.Args[2])
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <task-id> <new-description>")
			return
		}
		updateTask(parseID(os.Args[2]), os.Args[3])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <task-id>")
			return
		}
		deleteTask(parseID(os.Args[2]))
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <task-id>")
			return
		}
		markInProgress(parseID(os.Args[2]))
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <task-id>")
			return
		}
		markDone(parseID(os.Args[2]))
	case "list":
		var status string
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		listTasks(status)
	default:
		fmt.Println("Unknown command:", command)
	}
}

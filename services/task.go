package services

import (
	"fmt"
	"strconv"
	"strings"
	"taskmanager/models"
	"taskmanager/storage"
)

var Task_Json []models.Task_Json

func AddTask(title string) error {
	id := len(Task_Json) + 1
	task := models.Task_Json{ID: id, Title: title, Status: "pending"}
	Task_Json = append(Task_Json, task)
	if err := storage.SaveTasks(Task_Json); err != nil {
		return fmt.Errorf("failed to save tasks after adding: %w", err)
	}

	fmt.Printf("Task added: %d %s (%s)\n", task.ID, task.Title, task.Status)
	return nil
}

func ListTasks() error {
	if len(Task_Json) == 0 {
		fmt.Println("No tasks available.")
		return nil
	}
	for _, t := range Task_Json {
		fmt.Printf("%d %s (%s)\n", t.ID, t.Title, t.Status)
	}
	return nil
}

func UpdateTask(idStr string, status string) error {
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		return fmt.Errorf("invalid task id: %w", err)
	}
	for i, t := range Task_Json {
		if t.ID == id {
			Task_Json[i].Status = status

			if err := storage.SaveTasks(Task_Json); err != nil {
				return fmt.Errorf("failed to save tasks after update: %w", err)
			}

			fmt.Printf("Task updated: %d %s (%s)\n", t.ID, t.Title, t.Status)
			return nil
		}
	}

	return fmt.Errorf("task not found with id=%d", id)
}

func DeleteTask(idStr string) error {
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		return fmt.Errorf("invalid task id: %w", err)
	}
	for i, t := range Task_Json {
		if t.ID == id {
			Task_Json = append(Task_Json[:i], Task_Json[i+1:]...)

			if err := storage.SaveTasks(Task_Json); err != nil {
				return fmt.Errorf("failed to save tasks after delete: %w", err)
			}

			fmt.Printf("Task deleted: %d %s\n", t.ID, t.Title)
			return nil
		}
	}

	return fmt.Errorf("task not found with id=%d", id)
}

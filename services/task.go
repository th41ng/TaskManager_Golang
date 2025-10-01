package services

import (
	"fmt"
	"strconv"
	"taskmanager/models"
	"taskmanager/storage"
)

var Tasks []models.Task

func AddTask(title string) {
	id := len(Tasks) + 1
	task := models.Task_Json{ID: id, Title: title, Status: "pending"}
	Tasks = append(Tasks, task)
	storage.SaveTasks(Tasks)
	fmt.Printf("Task added: %d %s (%s)\n", task.ID, task.Title, task.Status)

}

func ListTasks() {
	if len(Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, t := range Tasks {
		fmt.Printf("%d %s (%s)\n", t.ID, t.Title, t.Status)
	}
}

func UpdateTask(idStr string, status string) {
	id, _ := strconv.Atoi(idStr)
	for i, t := range Tasks {
		if t.ID == id {
			Tasks[i].Status = status
			storage.SaveTasks(Tasks)
			fmt.Printf("Task updated: %d %s (%s)\n", t.ID, t.Title, t.Status)
			return
		}
	}
	fmt.Println("Task not found.")
}

func DeleteTask(idStr string) {
	id, _ := strconv.Atoi(idStr)
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			storage.SaveTasks(Tasks)
			fmt.Printf("Task deleted: %d %s\n", t.ID, t.Title)
			return
		}
	}
	fmt.Println("Task not found.")
}

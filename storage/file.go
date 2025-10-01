package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"taskmanager/models"
)

var FileName = "tasks.json"

func LoadTasks() []models.Task_Json {
	var tasks []models.Task_Json
	file, err := os.ReadFile(FileName)
	if err == nil {
		err = json.Unmarshal(file, &tasks)
		if err != nil {
			fmt.Print("File json lỗi")
		}
	}
	return tasks
}

func SaveTasks(tasks []models.Task_Json) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	err := os.WriteFile(FileName, data, 0644)
	if err != nil {
		fmt.Print("Lỗi")
	}
}

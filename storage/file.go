// package storage

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"taskmanager/models"
// )

// var FileName = "tasks.json"

// func LoadTasks() []models.Task_Json {
// 	var tasks []models.Task_Json
// 	file, err := os.ReadFile(FileName)
// 	if err == nil {
// 		err = json.Unmarshal(file, &tasks)
// 		if err != nil {
// 			fmt.Print("File json lỗi")
// 		}
// 	}
// 	return tasks
// }

//	func SaveTasks(tasks []models.Task_Json) {
//		data, _ := json.MarshalIndent(tasks, "", "  ")
//		err := os.WriteFile(FileName, data, 0644)
//		if err != nil {
//			fmt.Print("Lỗi")
//		}
//	}
package storage

import (
	"encoding/json"
	"os"
	"taskmanager/models"
)

var FileName = "tasks.json"

func LoadTasks() ([]models.Task_Json, error) {
	var tasks []models.Task_Json
	file, err := os.ReadFile(FileName)
	if err != nil {

		return tasks, nil
	}

	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []models.Task_Json) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(FileName, data, 0644); err != nil {
		return err
	}

	return nil
}

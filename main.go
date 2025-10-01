package main

import (
	// "bufio"
	// "fmt"
	// "os"
	// "strconv"
	// "taskmanager/services"
	// "taskmanager/storage"
	"taskmanager/database"
)

func main() {
	database.ConnectDatabase()
	// services.Tasks = storage.LoadTasks()

	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Println("1. Thêm task")
	// 	fmt.Println("2. Xem danh sách task")
	// 	fmt.Println("3. Đánh dấu task hoàn thành")
	// 	fmt.Println("4. Xóa task")
	// 	fmt.Println("5. Thoát")
	// 	fmt.Print("Chọn chức năng: ")

	// 	choiceStr, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Print("Lỗi")
	// 	}
	// 	choice, err := strconv.Atoi(choiceStr)
	// 	if err != nil {
	// 		fmt.Print("Lỗi")
	// 	}

	// 	switch choice {
	// 	case 1:
	// 		fmt.Print("Nhập tên task: ")
	// 		title, _ := reader.ReadString('\n')
	// 		services.AddTask(title)

	// 	case 2:
	// 		services.ListTasks()

	// 	case 3:
	// 		fmt.Print("Nhập ID task cần hoàn thành: ")
	// 		idStr, _ := reader.ReadString('\n')
	// 		services.UpdateTask(idStr, "done")

	// 	case 4:
	// 		fmt.Print("Nhập ID task cần xóa: ")
	// 		idStr, _ := reader.ReadString('\n')
	// 		services.DeleteTask(idStr)

	// 	case 5:
	// 		return

	// 	default:
	// 		fmt.Println("Lựa chọn không hợp lệ")
	// 	}
	// }
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"taskmanager/services"
	"taskmanager/storage"
)

func main() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		log.Fatalf("Không thể load tasks: %v", err)
	}
	services.Task_Json = tasks

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Thêm task")
		fmt.Println("2. Xem danh sách task")
		fmt.Println("3. Đánh dấu task hoàn thành")
		fmt.Println("4. Xóa task")
		fmt.Println("5. Thoát")
		fmt.Print("Chọn chức năng: ")

		choiceStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(fmt.Errorf("lỗi đọc input: %w", err))
			continue
		}

		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Vui lòng nhập số hợp lệ")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Nhập tên task: ")
			title, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(fmt.Errorf("không thể đọc tên task: %w", err))
				continue
			}
			title = strings.TrimSpace(title)

			if err := services.AddTask(title); err != nil {
				fmt.Println(err)
			}

		case 2:
			if err := services.ListTasks(); err != nil {
				fmt.Println(err)
			}

		case 3:
			fmt.Print("Nhập ID task cần hoàn thành: ")
			idStr, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(fmt.Errorf("không thể đọc ID: %w", err))
				continue
			}
			if err := services.UpdateTask(strings.TrimSpace(idStr), "done"); err != nil {
				// Check lỗi nhập ID không hợp lệ
				if errors.Is(err, strconv.ErrSyntax) {
					fmt.Println("Lỗi: ID nhập vào không phải số hợp lệ")
				} else {
					fmt.Println(err)
				}
			}

		case 4:
			fmt.Print("Nhập ID task cần xóa: ")
			idStr, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(fmt.Errorf("không thể đọc ID: %w", err))
				continue
			}
			if err := services.DeleteTask(strings.TrimSpace(idStr)); err != nil {
				if errors.Is(err, strconv.ErrSyntax) {
					fmt.Println("Lỗi: ID nhập vào không phải số hợp lệ")
				} else {
					fmt.Println(err)
				}
			}

		case 5:
			return

		default:
			fmt.Println("Lựa chọn không hợp lệ")
		}
	}
}

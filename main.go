package main

import (
	"context"
	"fmt"
	"log"
	"taskmanager/ent"
	"taskmanager/ent/task"
	"taskmanager/ent/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Kết nối database
	client, err := ent.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Auto migration
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// // Tạo User
	// u, err := client.User.
	// 	Create().
	// 	SetUsername("thang").
	// 	SetPassword("123").
	// 	Save(ctx)
	// if err != nil {
	// 	log.Fatalf("failed creating user: %v", err)
	// }

	// //  2 Project cho user
	// projs, err := client.Project.CreateBulk(
	// 	client.Project.Create().SetName("Đi học").SetOwner(u),
	// 	client.Project.Create().SetName("Đi làm").SetOwner(u),
	// ).Save(ctx)
	// if err != nil {
	// 	log.Fatalf("failed creating projects: %v", err)
	// }

	// // Tạo 3 Task cho mỗi Project
	// var allTasks []*ent.Task
	// for _, p := range projs {
	// 	tasks, err := client.Task.CreateBulk(
	// 		client.Task.Create().SetTitle("Task 1 - "+p.Name).SetProject(p),
	// 		client.Task.Create().SetTitle("Task 2 - "+p.Name).SetProject(p),
	// 		client.Task.Create().SetTitle("Task 3 - "+p.Name).SetProject(p),
	// 	).Save(ctx)
	// 	if err != nil {
	// 		log.Fatalf("failed creating tasks: %v", err)
	// 	}
	// 	allTasks = append(allTasks, tasks...)
	// }

	//Query user với project và task
	userWithProjects, err := client.User.
		Query().
		Where(user.ID(1)).
		WithProjects(func(pq *ent.ProjectQuery) {
			pq.WithTasks(func(tq *ent.TaskQuery) {
				tq.Where(task.DoneEQ(false))
			})
		}).
		Only(ctx)
	if err != nil {
		log.Fatalf("querry lỗi: %v", err)
	}

	fmt.Println("User:", userWithProjects.Username)
	for _, p := range userWithProjects.Edges.Projects {
		fmt.Println(" Project:", p.Name)
		for _, t := range p.Edges.Tasks {
			fmt.Printf("Task: %s, Done? %v\n", t.Title, t.Done)
		}
	}
}

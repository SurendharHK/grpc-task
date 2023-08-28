package main

import (
	"context"
	"fmt"
	pb "grpc-task/task"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	task := &pb.Task{Title: "Buy groceries",}
	addResp,err:=client.AddTask(context.Background(),task)
	if err!=nil{
		log.Fatal("Failed to add task: %v",err)
	}

	fmt.Printf("Added task with ID: %s\n",addResp.Id)
	tasksResp, err := client.GetTasks(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatal("Failed to retrieve tasks: %v", err)
	}
	fmt.Println("Tasks:")
	for _,task :=range tasksResp.Tasks{
		fmt.Printf("ID: %s,Title: %s,Completed: %v\n",task.Id,task.Title,task.Completed)

	}
}

package main

import (
	"context"
	"fmt"
	pb "grpc-task/task"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type taskServiceServer struct {
	mu    sync.Mutex
	tasks map[string]*pb.Task
	pb.UnimplementedTaskServiceServer
}

func (s *taskServiceServer) AddTask(ctx context.Context, req *pb.Task) (*pb.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	taskId := generateID()
	req.Id = taskId
	s.tasks[taskId] = req
	return &pb.TaskResponse{Id: taskId}, nil
}

func (s *taskServiceServer) GetTasks(ctx context.Context, req *pb.Empty) (*pb.TaskList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	tasks := make([]*pb.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return &pb.TaskList{Tasks: tasks}, nil
}

func generateID() string {
	return "taskID"
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("failed to listen :%v", err)
		return
	}
	server := grpc.NewServer()
	pb.RegisterTaskServiceServer(server, &taskServiceServer{
		tasks: make(map[string]*pb.Task),
	})
	fmt.Println("server listening on :50051")
	if err := server.Serve(lis); err != nil {
		fmt.Println("Failed to serve: %v", err)
	}
}

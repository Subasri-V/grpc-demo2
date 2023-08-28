package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	ts "grpc-demo2/tasks"

	"google.golang.org/grpc"
)

type server struct {
	mu    sync.Mutex
	tasks map[string]*ts.Task
	ts.UnimplementedTaskServiceServer
}

func (s *server) AddTask(ctx context.Context, req *ts.Task) (*ts.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	taskID := generateID()
	req.Id = taskID
	s.tasks[taskID] = req
	return &ts.TaskResponse{Id: taskID}, nil
}

func (s*server) GetTasks(ctx context.Context,req *ts.Empty)(*ts.TaskList,error){
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks:=make ([]*ts.Task,0,len(s.tasks))

	for _,task:=range s.tasks{
		tasks=append(tasks, task)
	}

	return &ts.TaskList{Tasks:tasks},nil
}

func generateID() string{
	return "1001"
}


func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to Listen: %v", err)
		return
	}

	s := grpc.NewServer()
	ts.RegisterTaskServiceServer(s, &server{
		tasks:make(map[string]*ts.Task),
	})

	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}

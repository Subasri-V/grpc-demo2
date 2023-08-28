package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	ts "grpc-demo2/tasks"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect : %v", err)
	}
	defer conn.Close()
	client:=ts.NewTaskServiceClient(conn)
	task:=&ts.Task{
		Title: "List of Books",
	}
	addresponse,err:=client.AddTask(context.Background(),task)


	if err != nil {
		log.Fatalf("failed to call Task : %v", err)
	}
	fmt.Printf("(Response)\nadded task with id  : %s\n", addresponse.Id)

	tasksResponse,err:=client.GetTasks(context.Background(),&ts.Empty{})
	if err!=nil{
		log.Fatalf("Failed to retrive tasks: %v",err)
	}
	fmt.Println("Tasks: ")
	for _,task:=range tasksResponse.Tasks{
		fmt.Printf("ID: %s, Title: %s, Completed: %v\n",task.Id,task.Title,task.Completed)
	}
}

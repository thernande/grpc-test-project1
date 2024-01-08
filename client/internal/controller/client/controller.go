package client

import (
	"context"
	"log"
	"time"

	v1 "github.com/thernande/grpc-test-project1/proto/todo/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AddTask(c v1.TodoServiceClient, description string, dueDate time.Time) uint64 {
	req := &v1.AddTaskRequest{
		Description: description,
		DueDate:     timestamppb.New(dueDate),
	}
	res, err := c.AddTask(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to add task: %v", err)
	}
	log.Printf("added task: %v", res)
	return res.Id
}

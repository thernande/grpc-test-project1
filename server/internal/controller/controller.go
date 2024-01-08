package controller

import (
	"context"

	v1 "github.com/thernande/grpc-test-project1/proto/todo/v1"
	"github.com/thernande/grpc-test-project1/server/internal/handler"
)

type Server struct {
	v1.UnimplementedTodoServiceServer
	d    handler.Db
	task []*v1.Task
}

func New(d handler.Db) *Server {
	return &Server{d: d}
}

func (s *Server) AddTask(c context.Context, req *v1.AddTaskRequest) (*v1.AddTaskResponse, error) {
	nextId := uint64(len(s.task) + 1)
	task := &v1.Task{
		Id:          nextId,
		Description: req.Description,
		DueDate:     req.DueDate,
	}

	s.task = append(s.task, task)
	return &v1.AddTaskResponse{Id: nextId}, nil

}

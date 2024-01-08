package memory

import (
	"time"

	v1 "github.com/thernande/grpc-test-project1/proto/todo/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MemoryDb struct {
	task []*v1.Task
}

func New() *MemoryDb {
	return &MemoryDb{task: []*v1.Task{}}
}

func (m *MemoryDb) AddTask(description string, dueDate time.Time) (uint64, error) {
	nextId := uint64(len(m.task) + 1)
	task := &v1.Task{
		Id:          nextId,
		Description: description,
		DueDate:     timestamppb.New(dueDate),
	}

	m.task = append(m.task, task)
	return nextId, nil

}

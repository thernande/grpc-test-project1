package handler

import "time"

type Db interface {
	AddTask(description string, dueDate time.Time) (uint64, error)
}

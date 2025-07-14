package storage

import (
	"sync"
	"time"
)

type IStorage interface{}

type TaskStatus string

var (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
	StatusFailed     TaskStatus = "failed"
)

type task struct {
	ID     uint
	Status TaskStatus
	Files  []string
	Errors []string

	ZipName string

	CreatedAt time.Time

	mu sync.Mutex
}

type storage struct {
	Tasks      map[string]*task
	ActiveTask int
	mu         sync.RWMutex
}

func NewStorage() IStorage {
	return &storage{
		Tasks: make(map[string]*task),
	}
}

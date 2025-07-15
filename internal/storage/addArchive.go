package storage

import (
	"errors"
	"fmt"
	"time"
)

// Создаем архив
func (s *storage) AddArchive() (*archive, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.ActiveArchives >= s.Cfg.ArchiveCount {
		return nil, errors.New("В данным момент сервер занят")
	}

	taskID := fmt.Sprintf("task_%d", time.Now().UnixNano())
	task := &archive{
		ID:        taskID,
		Status:    StatusPending,
		Files:     make([]string, 0),
		Errors:    make([]string, 0),
		CreatedAt: time.Now(),
	}

	s.Archives[taskID] = task
	s.ActiveArchives++
	return task, nil
}

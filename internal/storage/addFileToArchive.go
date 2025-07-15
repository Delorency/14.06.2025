package storage

import (
	"arch/internal"
	"errors"
	"path/filepath"
)

// Добавляем файл в архив
func (s *storage) AddFileToArchive(task *archive, path string) error {
	task.mu.Lock()
	defer task.mu.Unlock()

	filename := filepath.Base(path)
	if !internal.IsValidExtension(filename, s.Cfg.Extensions) {
		return errors.New("Недопустимый формат")
	}

	task.Files = append(task.Files, filename)

	return nil
}

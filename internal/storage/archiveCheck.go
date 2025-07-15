package storage

import (
	"errors"
)

// Проверяем архив
func (s *storage) ArchiveCheck(task *archive) error {
	if task.Status == StatusCompleted {
		return errors.New("Архив уже выполнен")
	}

	if len(task.Files) >= s.Cfg.ObjectCount {
		return errors.New("В архиве максимальное количество файлов")
	}

	return nil
}

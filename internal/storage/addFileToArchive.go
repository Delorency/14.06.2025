package storage

import (
	"arch/internal"
	"errors"
	"path/filepath"
)

// Добавляем файл в архив
func (s *storage) AddFileToArchive(arch *archive, path string) error {
	arch.mu.Lock()
	defer arch.mu.Unlock()

	if arch.Status != StatusPending {
		return errors.New("Архив не готов к добавлению файлов")
	}

	filename := filepath.Base(path)
	if !internal.IsValidExtension(filename, s.Cfg.Extensions) {
		return errors.New("Недопустимый формат")
	}

	arch.Files = append(arch.Files, path)

	if len(arch.Files) >= s.Cfg.ObjectCount {
		go s.ProcessOneArchive(arch)
	}

	return nil
}

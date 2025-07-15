package storage

import (
	"arch/internal"
	"fmt"
	"path/filepath"
)

func (s *storage) ProcessOneArchive(arch *archive) {

	arch.mu.Lock()

	arch.Status = StatusInProgress

	arch.mu.Unlock()

	files := make(map[string][]byte)
	for _, url := range arch.Files {
		data, err := internal.DownloadFile(url)
		if err != nil {
			arch.mu.Lock()
			arch.Errors = append(arch.Errors, fmt.Sprintf("Ошибка загрузки %s", url))
			arch.mu.Unlock()

			continue
		}
		filename := filepath.Base(url)
		files[filename] = data
	}

	zipBuffer, err := internal.CreateZipArchive(files)
	arch.mu.Lock()
	if err != nil {
		arch.Errors = append(arch.Errors, "Ошибка создания архива")
		arch.Status = StatusFailed

		arch.mu.Unlock()
	} else {
		arch.ZipBuffer = zipBuffer
		arch.ZipName = arch.ID + ".zip"
		arch.Status = StatusCompleted

		arch.mu.Unlock()
		s.CompleteArchive(arch.ID)
	}
}

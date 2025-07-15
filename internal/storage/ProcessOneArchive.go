package storage

import (
	"arch/internal"
	"fmt"
	"path/filepath"
)

func (s *storage) ProcessOneArchive(arch *archive) {
	arch.mu.Lock()
	defer arch.mu.Unlock()

	arch.Status = StatusInProgress

	files := make(map[string][]byte)
	for _, url := range arch.Files {
		data, err := internal.DownloadFile(url)
		if err != nil {
			arch.Errors = append(arch.Errors, fmt.Sprintf("Ошибка загрузки %s", url))

			continue
		}
		filename := filepath.Base(url)
		files[filename] = data
	}

	zipBuffer, err := internal.CreateZipArchive(files)
	if err != nil {
		arch.Errors = append(arch.Errors, "Ошибка создания архива")
		arch.Status = StatusFailed

	} else {
		arch.ZipBuffer = zipBuffer
		arch.ZipName = arch.ID + ".zip"
		arch.Status = StatusCompleted

		s.CompleteArchive(arch.ID)
	}
}

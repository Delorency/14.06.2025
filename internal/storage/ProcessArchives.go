package storage

import (
	"arch/internal"
	"fmt"
	"path/filepath"
	"time"
)

func (s *storage) ProcessArchive() {
	for {
		time.Sleep(1 * time.Second)

		s.mu.Lock()
		for _, arch := range s.Archives {
			if arch.Status == StatusPending && len(arch.Files) == s.Cfg.ObjectCount {

				go func(t *archive) {
					t.mu.Lock()
					t.Status = StatusInProgress
					t.mu.Unlock()

					files := make(map[string][]byte)
					for _, url := range t.Files {
						data, err := internal.DownloadFile(url)
						if err != nil {
							t.Errors = append(t.Errors, fmt.Sprintf("Ошибка загрузки %s", url))
							continue
						}
						filename := filepath.Base(url)
						files[filename] = data
					}

					if len(files) > 0 {
						zipBuffer, err := internal.CreateZipArchive(files)
						if err != nil {
							t.Errors = append(t.Errors, "Ошибка создания архива")
							t.Status = StatusFailed
						} else {
							t.ZipBuffer = zipBuffer
							t.ZipName = t.ID + ".zip"
							t.Status = StatusCompleted
						}
					} else {
						t.Status = StatusFailed
					}

					s.CompleteArchive(t.ID)
				}(arch)

			}
		}
		s.mu.Unlock()
	}
}

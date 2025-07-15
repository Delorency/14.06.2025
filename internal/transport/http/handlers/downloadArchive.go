package handlers

import (
	"arch/internal/storage"
	. "arch/internal/transport/http/error"
	"bytes"
	"io"
	"net/http"
	"strings"
)

func (ah *archiveHandler) DownloadArchive(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := strings.TrimPrefix(r.URL.Path, "/download/")
	taskID := strings.TrimSuffix(filename, ".zip")

	task, exists := ah.Storage.GetArchive(taskID)
	if !exists {
		Error(w, "Архив не найден", http.StatusNotFound)
		return
	}
	if task.Status != storage.StatusCompleted {
		Error(w, "Архив не выполнен", http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	if _, err := io.Copy(w, bytes.NewReader(task.ZipBuffer.Bytes())); err != nil {
		Error(w, "Ошибка отправки архиве", http.StatusInternalServerError)
	}
}

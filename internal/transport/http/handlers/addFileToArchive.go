package handlers

import (
	"arch/internal"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (ah *archiveHandler) AddFileToArchive(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodybytes, _ := io.ReadAll(r.Body)

	var req internal.AddFileToArchiveRequest

	json.Unmarshal(bodybytes, &req)
	if req.Url == "" {
		http.Error(w, "Укажите `url` в теле запроса", http.StatusBadRequest)
	}

	taskID := strings.TrimPrefix(r.URL.Path, "/add-file/")
	task, exists := ah.Storage.GetArchive(taskID)
	if !exists {
		http.Error(w, "Архив не найден", http.StatusNotFound)
		return
	}

	err := ah.Storage.ArchiveCheck(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = ah.Storage.AddFileToArchive(task, req.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

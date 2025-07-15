package handlers

import (
	"arch/internal"
	. "arch/internal/transport/http/error"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (ah *archiveHandler) AddFileToArchive(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodybytes, _ := io.ReadAll(r.Body)

	var req internal.AddFileToArchiveRequest

	json.Unmarshal(bodybytes, &req)
	if req.Url == "" {
		Error(w, "Укажите `url` в теле запроса", http.StatusBadRequest)
		return
	}

	taskID := strings.TrimPrefix(r.URL.Path, "/add-file/")
	task, exists := ah.Storage.GetArchive(taskID)
	if !exists {
		Error(w, "Архив не найден", http.StatusNotFound)
		return
	}

	err := ah.Storage.AddFileToArchive(task, req.Url)
	if err != nil {
		Error(w, err.Error(), http.StatusBadRequest)
	}
}

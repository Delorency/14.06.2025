package handlers

import (
	"arch/internal"
	"encoding/json"
	"net/http"
	"strings"
)

func (ah *archiveHandler) CheckArchiveStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskID := strings.TrimPrefix(r.URL.Path, "/status/")
	task, exists := ah.Storage.GetArchive(taskID)
	if !exists {
		http.Error(w, "Архив не найден", http.StatusNotFound)
		return
	}

	res := internal.CheckArchiveStatusResponse{
		TaskID:   taskID,
		Status:   string(task.Status),
		Errors:   task.Errors,
		Download: ah.Storage.AddDownload(task),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

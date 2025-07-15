package handlers

import (
	"arch/internal"
	"arch/internal/storage"
	. "arch/internal/transport/http/error"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (ah *archiveHandler) CheckArchiveStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskID := strings.TrimPrefix(r.URL.Path, "/status/")
	task, exists := ah.Storage.GetArchive(taskID)
	if !exists {
		Error(w, "Архив не найден", http.StatusNotFound)
		return
	}

	res := internal.CheckArchiveStatusResponse{
		TaskID:     taskID,
		Status:     string(task.Status),
		FilesCount: len(task.Files),
		Errors:     task.Errors,
	}
	if task.Status == storage.StatusCompleted {
		res.Download = fmt.Sprintf("/download/%s.zip", res.TaskID)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

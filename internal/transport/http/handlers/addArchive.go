package handlers

import (
	"arch/internal"
	. "arch/internal/transport/http/error"
	"encoding/json"
	"net/http"
)

func (ah *archiveHandler) AddArchive(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task, err := ah.Storage.AddArchive()
	if err != nil {
		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := internal.AddArchiveResponse{TaskID: task.ID, CreatedAt: task.CreatedAt}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

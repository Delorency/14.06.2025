package internal

import (
	"time"
)

type AddArchiveResponse struct {
	TaskID    string    `json:"taskID"`
	CreatedAt time.Time `json:"createdAt"`
}

type AddFileToArchiveRequest struct {
	Url string `json:"url"`
}

type CheckArchiveStatusResponse struct {
	TaskID   string   `json:"taskID"`
	Status   string   `json:"status"`
	Errors   []string `json:"errors"`
	Download string   `json:"download"`
}

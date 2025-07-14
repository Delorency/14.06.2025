package handlers

import "net/http"

type ITaskHandler interface {
	AddArchive(w http.ResponseWriter, r *http.Request)
	AddFileToArchive(w http.ResponseWriter, r *http.Request)
	CheckArchiveStatus(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct{}

func NewTaskHandler() ITaskHandler {
	return &taskHandler{}
}

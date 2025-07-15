package handlers

import (
	"arch/internal/storage"
	"net/http"
)

type IArchiveHandler interface {
	AddArchive(w http.ResponseWriter, r *http.Request)
	AddFileToArchive(w http.ResponseWriter, r *http.Request)
	CheckArchiveStatus(w http.ResponseWriter, r *http.Request)
	DownloadArchive(w http.ResponseWriter, r *http.Request)
}

type archiveHandler struct {
	Storage storage.IStorage
}

func NewTaskHandler(storage storage.IStorage) IArchiveHandler {
	return &archiveHandler{Storage: storage}
}

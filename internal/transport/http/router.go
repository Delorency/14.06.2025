package http

import (
	config "arch/internal"
	"arch/internal/storage"
	"arch/internal/transport/http/handlers"
	"net/http"
)

func NewRouter(cfgArch *config.ConfigArchive, storage storage.IStorage) *http.ServeMux {
	handler := handlers.NewTaskHandler(storage)

	router := http.NewServeMux()

	router.HandleFunc("/add-archive", handler.AddArchive)
	router.HandleFunc("/add-file/", handler.AddFileToArchive)
	router.HandleFunc("/status/", handler.CheckArchiveStatus)
	router.HandleFunc("/download/", handler.DownloadArchive)

	return router
}

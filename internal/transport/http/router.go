package http

import (
	"arch/internal/transport/http/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux {
	handler := handlers.NewTaskHandler()

	router := http.NewServeMux()

	router.HandleFunc("/add-archive", handler.AddArchive)
	router.HandleFunc("/add-file/", handler.AddFileToArchive)
	router.HandleFunc("/status/", handler.CheckArchiveStatus)

	return router
}

package http

import (
	"arch/internal"
	"arch/internal/storage"
	"fmt"
	"net/http"
)

func NewHTTPServer(cfghttp *internal.ConfigHTTPServer, cfgArch *internal.ConfigArchive, storage storage.IStorage) *http.Server {
	router := NewRouter(cfgArch, storage)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfghttp.Host, cfghttp.Port),
		Handler: router,
	}

	return server
}

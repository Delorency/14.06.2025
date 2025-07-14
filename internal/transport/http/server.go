package http

import (
	config "arch/internal"
	"fmt"
	"net/http"
)

func NewHTTPServer(cfghttp *config.ConfigHTTPServer) *http.Server {
	router := NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfghttp.Host, cfghttp.Port),
		Handler: router,
	}

	return server
}

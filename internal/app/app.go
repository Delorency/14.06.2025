package app

import (
	config "arch/internal"
	ht "arch/internal/transport/http"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	cfg := config.MustLoad()

	server := ht.NewHTTPServer(cfg.Http)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Сервер запущен на %s:%s\n", cfg.Http.Host, cfg.Http.Port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Ошибка работы сервера: %v\n", err)
		}
	}()

	<-ctx.Done()

	log.Println("Сервер завершает свою работу...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Ошибка грациозного завершения: %v\n", err)

		if err := server.Close(); err != nil {
			log.Fatalf("Ошибка принудительного завершения: %v\n", err)
		}
	}

}

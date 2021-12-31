package main

import (
	"context"
	"github.com/Kichiyaki/wasmplayground/assets"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Handler:           http.FileServer(http.FS(assets.FS)),
		Addr:              ":8080",
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       2 * time.Second,
	}

	go func(srv *http.Server) {
		ctx, stop := signal.NotifyContext(
			context.Background(),
			os.Interrupt,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
		defer stop()

		<-ctx.Done()

		log.Println("shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Fatal("srv.Shutdown:", err)
		}
	}(srv)

	log.Println("Server is listening on the port 8080")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("srv.ListenAndServe:", err)
	}
}

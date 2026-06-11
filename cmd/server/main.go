package main

import (
	"os"
	"log"
	"net/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})

	srv := http.Server{Addr: ":" + port, Handler: mux}

	log.Printf("Сервер запущен на порту %s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка запуска сервера!")
	}
}

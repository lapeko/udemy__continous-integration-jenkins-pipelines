package main

import (
	"fmt"
	"net/http"
	"os"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	return mux
}

func main() {
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	mux := SetupRouter()

	host := fmt.Sprintf("localhost:%s", PORT)
	fmt.Printf("Server running on %s\n", host)
	if err := http.ListenAndServe(host, mux); err != nil {
		fmt.Printf("Error: %e\n", err)
	}
}

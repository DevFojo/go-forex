package app

import (
	"fmt"
	"net/http"

	"github.com/mrfojo/go-forex/src/handlers"
)

func Run() {
	httpServer := &http.Server{
		Addr:    ":5000",
		Handler: handler(),
	}
	fmt.Println("Server running on port 5000")
	httpServer.ListenAndServe()
}

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/rates/", handlers.HandleRateRequests)
	return mux
}

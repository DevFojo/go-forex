package app

import (
	"fmt"
	"net/http"

	"github.com/devFojo/go-forex/handlers"
)

func Run() error {
	httpServer := &http.Server{
		Addr:    ":5000",
		Handler: Handler,
	}
	fmt.Println("Server running on port 5000")
	return httpServer.ListenAndServe()
}

var Handler *http.ServeMux

func init() {
	Handler = http.NewServeMux()
	Handler.HandleFunc("/rates/", handlers.HandleRateRequests)
}

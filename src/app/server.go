package app

import (
	"fmt"
	"net/http"

	"github.com/mrfojo/go-forex/src/handlers"
)

func Run() {
	httpServer := &http.Server{
		Addr:    ":5000",
		Handler: Handler,
	}
	fmt.Println("Server running on port 5000")
	httpServer.ListenAndServe()
}

var Handler *http.ServeMux

func init() {
	Handler = http.NewServeMux()
	Handler.HandleFunc("/rates/", handlers.HandleRateRequests)
}

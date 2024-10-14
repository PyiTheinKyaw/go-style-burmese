package main

import (
	"fmt"
	"net/http"
)

// Define the Handler struct
type Handler struct{}

// Implement the ServeHTTP method required by http.Handler Interface
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

// Interface compliance check at compile-time
var _ http.Handler = (*Handler)(nil)

func main() {
	var h http.Handler = &Handler{} // Valid because *Handler implements http.Handler
	http.ListenAndServe(":8080", h)
}

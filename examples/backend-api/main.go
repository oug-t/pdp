package main

import (
	"fmt"
	"net/http"
	"pdp/libs/go" // Assuming local module setup
)

func main() {
	mux := http.NewServeMux()

	// Wrap your handler with PDP middleware
	// Defaulting to Level 0 (Private) if header is missing
	handler := pdp.Middleware(0)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level, _ := pdp.Level(r.Context())
		fmt.Fprintf(w, "Request processed with PDP Level: %d", level)
	}))

	mux.Handle("/api/chat", handler)
	fmt.Println("PDP Demo Server running on :8080")
	http.ListenAndServe(":8080", mux)
}

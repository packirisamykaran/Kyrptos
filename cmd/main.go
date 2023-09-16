package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/packirisamykaran/kryptos/internal/request/routes"
)

func main() {
	// This handler registers a response for the root path "/"
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Kryptos Backend is running!")
	// })

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Connect all the routes in routes.go to the server
	routes.ProtocolWhalesRoute(r)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kryptos Backend is running!")
	}).Methods("GET")
	// Use the Gorilla Mux router as the main router
	http.Handle("/", r)

	// Set up the server address
	port := "8000" // You mentioned using port 8000
	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", serverAddr)

	// Start the HTTP server
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Printf("Server error: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/packirisamykaran/kryptos/internal/request/routes"
)

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define the root route ("/") and set it to the "index" function
	router.HandleFunc("/", index)

	// Connect all the routes in routes.go to the server under the "/whale" path
	mount(router, "/whale", routes.ProtocolWhalesRoutes())

	// Set up the server address and port
	port := "8000"
	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", serverAddr)

	// Start the HTTP server and handle requests with the Gorilla Mux router
	log.Fatal(http.ListenAndServe(serverAddr, router))
}

// mount is a helper function to mount a sub-router under a specific path
func mount(r *mux.Router, path string, handler http.Handler) {
	// Use PathPrefix to match all paths under the specified "path"
	// and use StripPrefix to remove the specified "path" prefix from the URL
	r.PathPrefix(path).Handler(http.StripPrefix(path, handler))
}

// index is the handler for the root ("/") route
func index(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message
	fmt.Fprintf(w, "Kryptos Backend is running!")
}

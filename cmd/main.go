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
	router.HandleFunc("/", index)

	// Connect all the routes in routes.go to the server
	mount(router, "/whale", routes.ProtocolWhalesRoutes())

	// Use the Gorilla Mux router as the main router

	// Set up the server address
	port := "8000"
	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))

}
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(http.StripPrefix(path, handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Kryptos Backend is running!")
}

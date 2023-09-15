package main

import (
	"fmt"
	"net/http"

	"github.com/packirisamykaran/kryptos/internal/services/whale"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kryptos Backend is running!")
	})

	whale.GetWhalesByProtocol("Solend")

	port := "8000" // You mentioned using port 8000

	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", serverAddr)

	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Printf("Server error: %v", err)
	}
}

package routes

import (
	"github.com/gorilla/mux"
	"github.com/packirisamykaran/kryptos/internal/request/handlers/whale"
)

func ProtocolWhalesRoute(r *mux.Router) {
	r.HandleFunc("/getWhalesByProtocol", whale.GetWhalesByProtocol).Methods("GET")
}

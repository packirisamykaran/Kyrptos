package routes

import (
	"github.com/gorilla/mux"
	"github.com/packirisamykaran/kryptos/internal/request/handlers/whale"
)

func ProtocolWhalesRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getWhalesByProtocol", whale.GetWhalesByProtocol).Methods("GET")
	return router
}

package whale

import (
	"encoding/json"
	"net/http"

	"github.com/packirisamykaran/kryptos/internal/models"
	"github.com/packirisamykaran/kryptos/internal/services/whale"
	"github.com/packirisamykaran/kryptos/internal/utils"
)

type Whale = models.Whale

type WhaleByProtocolResponse struct {
	Whales []Whale `json:"whales"`
	Error  error   `json:"error"`
}

func GetWhalesByProtocol(w http.ResponseWriter, r *http.Request) {

	var (
		whales      []Whale
		err         error
		protocol    string
		reponseBody WhaleByProtocolResponse
	)

	protocol = r.URL.Query().Get("protocol")

	whales, err = whale.GetWhalesByProtocol(protocol)
	if err != nil {
		reponseBody.Whales = whales
		reponseBody.Error = utils.HandleError(err)
		body, _ := json.Marshal(reponseBody)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		return
	}

	reponseBody.Whales = whales
	reponseBody.Error = err
	body, _ := json.Marshal(reponseBody)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

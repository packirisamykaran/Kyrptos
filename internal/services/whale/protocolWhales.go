package whale

import (
	"encoding/json"
	"strconv"

	"github.com/packirisamykaran/kryptos/internal/models"
	"github.com/packirisamykaran/kryptos/internal/utils"
)

type Whale = models.Whale

type WhaleByProtocolResponse struct {
	Data  []Whale `json:"data"`
	Total int     `json:"total"`
}

func GetWhalesByProtocol(protocol string) (whales []Whale, err error) {
	var limit = 10
	var offset = 0
	var protoclAddress = DefiProtocolAddressMap[protocol]

	params := map[string]string{
		"tokenAddress": protoclAddress,
		"limit":        strconv.Itoa(limit),
		"offset":       strconv.Itoa(offset),
	}

	resBodyByte, err := utils.GetRequestWithParams(SolscanAPI, params)
	if err != nil {
		return whales, err
	}

	var responseBody WhaleByProtocolResponse

	err = json.Unmarshal(resBodyByte, &responseBody)
	if err != nil {
		return whales, err
	}

	whales = responseBody.Data

	return whales, nil

}

package whale

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/packirisamykaran/kryptos/internal/models"
	"github.com/packirisamykaran/kryptos/internal/utils"
)

type Whale = models.Whale

type WhaleByProtocolResponse struct {
	Data  []Whale `json:"data"`
	Total int     `json:"total"`
}

func GetWhalesByProtocol(protocol string) []Whale {
	var limit = 10
	var offset = 0
	var protoclAddress = DefiProtocolAddressMap[protocol]
	// var query = fmt.Sprintf("%stokenAddress=%s&limit=%d&offset=%d", SolscanAPI, DefiProtocolAddressMap[protocol], limit, offset)

	params := map[string]string{
		"tokenAddress": protoclAddress,
		"limit":        strconv.Itoa(limit),
		"offset":       strconv.Itoa(offset),
	}

	resBodyByte, err := utils.GetRequestWithParams(SolscanAPI, params)
	if err != nil {
		fmt.Println("Error:", err)

	}

	var responseBody WhaleByProtocolResponse
	var whales []Whale

	err = json.Unmarshal(resBodyByte, &responseBody)
	if err != nil {
		fmt.Println(err)
	}

	whales = responseBody.Data

	return whales

}

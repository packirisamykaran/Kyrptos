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

func GetWhalesByProtocol(protocol string) (whales []Whale, err error) {
	var limit = 10
	var offset = 0
	var protoclAddress = DefiProtocolAddressMap[protocol]

	params := map[string]string{
		"tokenAddress": protoclAddress,
		"limit":        strconv.Itoa(limit),
		"offset":       strconv.Itoa(offset),
	}

	resBodyByte, err := utils.GetRequestWithParams(SolscanTokenHolderAPI, params)
	if err != nil {
		return whales, err
	}

	var responseBody WhaleByProtocolResponse

	err = json.Unmarshal(resBodyByte, &responseBody)
	if err != nil {
		return whales, err
	}

	whales = responseBody.Data

	_, errr := getTokensByOwner(whales[0].Owner)
	if errr != nil {
		fmt.Println(utils.HandleError(errr))
	}

	return whales, nil

}

func GetAllProtocolWhales() (protocolWhalesMap map[string][]Whale, err error) {

	for protocol, address := range DefiProtocolAddressMap {

		whales, err := GetWhalesByProtocol(address)
		if err != nil {
			return protocolWhalesMap, err
		}

		protocolWhalesMap[protocol] = whales

	}

	return protocolWhalesMap, nil

}

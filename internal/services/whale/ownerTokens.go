package whale

import (
	"encoding/json"
	"fmt"

	"github.com/packirisamykaran/kryptos/internal/models"
	"github.com/packirisamykaran/kryptos/internal/utils"
)

type TokenHeld = models.TokenHeld

// Get Tokens held by a whale using the whale's ownder address
func getTokensByOwner(owner string) (tokens []TokenHeld, err error) {

	params := map[string]string{
		"account": owner,
	}

	resBodyByte, err := utils.GetRequestWithParams(SolscanAccountTokensAPI, params)
	if err != nil {
		return tokens, err
	}

	var responseBody []TokenHeld

	err = json.Unmarshal(resBodyByte, &responseBody)
	if err != nil {
		return tokens, err
	}

	tokens = responseBody

	fmt.Println(tokens)

	return tokens, nil

}

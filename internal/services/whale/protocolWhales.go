package whale

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetProtocolWhaleByAddress(protocol string) {
	var limit = 10
	var offset = 0

	var query = fmt.Sprintf("%stokenAddress=%s&limit=%d&offset=%d", SolscanAPI, DefiProtocolAddressMap[protocol], limit, offset)

	res, err := http.Get(query)
	if err != nil {
		fmt.Println(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(resBody))
}

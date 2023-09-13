package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kryptos Backend is running!")
	})

	port := "8000" // You mentioned using port 8000

	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", serverAddr)

	GetTop20Holder("SLNDpmoWTVADgEdndyvWzroNL7zSi1dF9PC3xHGtPwp")

	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Printf("Server error: %v", err)
	}
}

func GetTop20Holder(protocol string) {

	var API = "https://public-api.solscan.io/token/holders?"
	var limit = 10
	var offset = 0

	var query = fmt.Sprintf("%stokenAddress=%s&limit=%d&offset=%d", API, protocol, limit, offset)

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

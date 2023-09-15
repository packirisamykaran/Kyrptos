package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetWithParams sends a GET request to the specified URL with the provided query parameters.
func GetRequestWithParams(api string, params map[string]string) ([]byte, error) {
	// Parse URL
	u, err := url.Parse(api)
	if err != nil {
		return nil, err
	}

	// Add query parameters to the URL
	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	// Send the GET request
	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

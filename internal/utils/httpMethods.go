package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetWithParams sends a GET request to the specified URL with the provided query parameters.
func GetRequestWithParams(api string, params map[string]string) ([]byte, error) {
	// Parse the API URL
	parsedURL, err := url.Parse(api)
	if err != nil {
		return nil, err
	}

	// Add query parameters to the URL
	queryParameters := parsedURL.Query()
	for key, value := range params {
		queryParameters.Add(key, value)
	}
	parsedURL.RawQuery = queryParameters.Encode()

	// Send the GET request
	response, err := http.Get(parsedURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

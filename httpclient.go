package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get invokes GET request aand returns response body
func Get(url string) ([]byte, error) {
	return execute("GET", url)
}

func execute(method string, url string) ([]byte, error) {

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, fmt.Errorf("GET %s failed: %s", url, response.Status)
	}

	if response.Body == nil {
		return []byte{}, nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

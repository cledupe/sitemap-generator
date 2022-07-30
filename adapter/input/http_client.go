package input

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	httpClient := &HttpClient{client: new(http.Client)}
	httpClient.client.Timeout = 5 * time.Second
	return httpClient
}

func (httpClient HttpClient) GetData(url string) (string, error) {

	if url != "" {
		resp, err := httpClient.client.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if resp.Status != "200 OK" {
			return "", fmt.Errorf("response status difrente of 200")
		}

		content, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}

		return string(content), nil

	}
	return "", errors.New("No path found")
}

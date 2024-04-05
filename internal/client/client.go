package client

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

const (
	httpTimeout = 2 * time.Second
)

func Healthcheck(endpoint string) bool {
	client := &http.Client{Timeout: httpTimeout}

	url := fmt.Sprintf("%s/healthcheck", endpoint)

	resp, err := client.Get(url)

	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}

		return false
	}

	return true
}

func Get(url string) (*http.Response, error) {
	client := &http.Client{Timeout: httpTimeout}

	resp, err := client.Get(url)

	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %w", url, err)
	}

	defer resp.Body.Close()

	return resp, nil
}

func Post(endpoint, path, contentType string, payload []byte) (*http.Response, error) {
	httpClient := &http.Client{Timeout: httpTimeout}

	url := fmt.Sprintf("%s/%s", endpoint, path)

	resp, err := httpClient.Post(url, contentType, bytes.NewBuffer(payload)) /* #nosec G107 */
	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %w", url, err)
	}

	defer resp.Body.Close()

	return resp, nil
}

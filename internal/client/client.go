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

type HTTPClient struct {
	Client *http.Client
}

func New() *HTTPClient {
	return &HTTPClient{
		Client: &http.Client{Timeout: httpTimeout},
	}
}

func (c *HTTPClient) Healthcheck(endpoint string) bool {
	url := fmt.Sprintf("%s/healthcheck", endpoint)

	resp, err := c.Get(url)

	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}

		return false
	}

	return true
}

func (c *HTTPClient) Get(endpoint string) (*http.Response, error) {
	resp, err := c.Client.Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("get request to %s failed: %w", endpoint, err)
	}

	defer resp.Body.Close()

	return resp, nil
}

func (c *HTTPClient) Post(endpoint, contentType string, payload []byte) error {
	resp, err := c.Client.Post(endpoint, contentType, bytes.NewBuffer(payload)) /* #nosec G107 */
	if err != nil {
		return fmt.Errorf("post request to %s failed: %w", endpoint, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("post request returned status %d", resp.StatusCode)
	}

	return nil
}

func (c *HTTPClient) Delete(endpoint string) error {
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("delete request to %s failed: %w", endpoint, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete request returned status %d", resp.StatusCode)
	}

	return nil
}

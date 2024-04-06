package origin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type (
	ReporterResponse struct {
		Signal       string   `json:"signal"`
		Packagers    []string `json:"packagers"`
		LastReported string   `json:"last_reported"`
	}

	SignalResponse struct {
		Signal string `json:"signal"`
		Server string `json:"server"`
	}
)

func getSignalPackagers(reporterEndpoint, signal string) (*ReporterResponse, error) {
	client := &http.Client{Timeout: 2 * time.Second}

	endpoint := fmt.Sprintf("%s/ingest/%s", reporterEndpoint, signal)

	resp, err := client.Get(endpoint)
	if err != nil {
		return &ReporterResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := errors.New("there is no signal active")
		return &ReporterResponse{}, errMsg
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ReporterResponse{}, err
	}

	var response ReporterResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return &ReporterResponse{}, err
	}

	return &response, nil
}

func formatPath(packagers []string, signal string) SignalResponse {
	path := fmt.Sprintf("%s/%s/playlist.m3u8", packagers[0], signal)
	activeSignalPath := SignalResponse{
		Signal: signal,
		Server: path,
	}

	return activeSignalPath
}

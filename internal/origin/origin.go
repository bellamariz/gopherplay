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

func listSignals(reporterEndpoint string) ([]string, error) {
	client := &http.Client{Timeout: 2 * time.Second}

	endpoint := fmt.Sprintf("%s/ingests", reporterEndpoint)

	resp, err := client.Get(endpoint)
	if err != nil {
		return []string{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := errors.New("there is no active signal")
		return []string{}, errMsg
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	var response []ReporterResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return []string{}, err
	}

	signals := []string{}
	for _, v := range response {
		signals = append(signals, v.Signal)
	}

	return signals, nil
}

func getSignalPackagers(reporterEndpoint, signal string) (*ReporterResponse, error) {
	client := &http.Client{Timeout: 2 * time.Second}

	endpoint := fmt.Sprintf("%s/ingest/%s", reporterEndpoint, signal)

	resp, err := client.Get(endpoint)
	if err != nil {
		return &ReporterResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := errors.New("there is no active signal")
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

package origin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bellamariz/go-live-without-downtime/internal/sources"
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

	var response []sources.Ingest
	if err := json.Unmarshal(body, &response); err != nil {
		return []string{}, err
	}

	signals := []string{}
	for _, v := range response {
		signals = append(signals, v.Signal)
	}

	return signals, nil
}

func getSignalPackagers(reporterEndpoint, signal string) (*sources.Ingest, error) {
	client := &http.Client{Timeout: 2 * time.Second}

	endpoint := fmt.Sprintf("%s/ingest/%s", reporterEndpoint, signal)

	resp, err := client.Get(endpoint)
	if err != nil {
		return &sources.Ingest{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := errors.New("there is no active signal")
		return &sources.Ingest{}, errMsg
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &sources.Ingest{}, err
	}

	var response sources.Ingest
	if err := json.Unmarshal(body, &response); err != nil {
		return &sources.Ingest{}, err
	}

	return &response, nil
}

func formatPath(packagers []string, signal string) sources.Source {
	path := fmt.Sprintf("%s/%s/playlist.m3u8", packagers[0], signal)
	activeSignalPath := sources.Source{
		Signal: signal,
		Server: path,
	}

	return activeSignalPath
}

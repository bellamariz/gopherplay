package origin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bellamariz/go-live-without-downtime/internal/client"
	"github.com/bellamariz/go-live-without-downtime/internal/sources"
)

func listSignals(reporterEndpoint string) ([]string, error) {
	httpClient := client.New()

	endpoint := fmt.Sprintf("%s/ingests", reporterEndpoint)

	resp, err := httpClient.Get(endpoint)
	if err != nil {
		return []string{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []string{}, errors.New("no active signals found by origin")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	defer resp.Body.Close()

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

func getSignalIngest(reporterEndpoint, signal string) (*sources.Ingest, error) {
	httpClient := client.New()

	endpoint := fmt.Sprintf("%s/ingests/%s", reporterEndpoint, signal)

	resp, err := httpClient.Get(endpoint)
	if err != nil {
		return &sources.Ingest{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return &sources.Ingest{}, errors.New("no active signals found by origin")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &sources.Ingest{}, err
	}

	defer resp.Body.Close()

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

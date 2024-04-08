package sources

import (
	"time"

	"github.com/rs/zerolog/log"
)

type Ingest struct {
	Signal       string   `json:"signal"`
	Packagers    []string `json:"packagers"`
	LastReported string   `json:"last_reported"`
}

type Source struct {
	Signal string `json:"signal"`
	Server string `json:"server"`
}

func (i Ingest) IsActive() bool {
	t, err := time.Parse(time.RFC1123, i.LastReported)
	if err != nil {
		log.Warn().Err(err).Msg("Error when verifying if ingest is active")
		return false
	}

	wasRecentlyReported := time.Since(t) <= (15 * time.Second)
	hasActiveIngest := len(i.Packagers) > 0

	return wasRecentlyReported && hasActiveIngest
}

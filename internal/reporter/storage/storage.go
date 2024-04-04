package storage

import (
	"github.com/rs/zerolog/log"
	"github.com/sopherapps/go-scdb/scdb"
)

type Storage struct {
	Store *scdb.Store
}

var (
	maxKeys            uint64 = 1_000_000
	redundantBlocks    uint16 = 1
	poolCapacity       uint64 = 10
	compactionInterval uint32 = 1_800
)

func New() *Storage {
	store, err := scdb.New(
		"db",
		&maxKeys,
		&redundantBlocks,
		&poolCapacity,
		&compactionInterval,
		true,
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create storage")
	}

	defer func() {
		_ = store.Close()
	}()

	return &Storage{
		Store: store,
	}
}

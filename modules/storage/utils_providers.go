package storage

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// UpdateProviders gets the updated providers list and stores it inside the database
func (m *Module) UpdateProviders(height int64) error {
	log.Debug().Str("module", "storage").Int64("height", height).
		Msg("updating providers")

	providers, err := m.source.GetProviders(height)
	if err != nil {
		return fmt.Errorf("error while getting providers list: %s", err)
	}

	return m.db.SaveStorageProviders(providers, height)
}

// UpdateActiveProviders gets the updated ACTIVE providers list and stores it inside the database
func (m *Module) UpdateActiveProviders(height int64) error {
	log.Debug().Str("module", "storage").Int64("height", height).
		Msg("updating ACTIVE providers")

	activeProviders, err := m.source.GetActiveProviders(height)
	if err != nil {
		return fmt.Errorf("error while getting providers list: %s", err)
	}

	return m.db.SaveActiveProviders(activeProviders, height)
}

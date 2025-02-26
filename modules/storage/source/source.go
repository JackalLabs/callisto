package source

import (
	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

type Source interface {
	GetParams(height int64) (storagetypes.Params, error)
	GetProviders(height int64) ([]storagetypes.Providers, error)
	GetFiles(height int64) ([]storagetypes.UnifiedFile, error)
	GetActiveProviders(height int64) ([]storagetypes.ActiveProviders, error)
}

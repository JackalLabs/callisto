package source

import (
	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

type Source interface {
	Params(height int64) (storagetypes.Params, error)
	Providers(height int64) ([]storagetypes.Providers, error)
	Files(height int64) ([]storagetypes.UnifiedFile, error)
}

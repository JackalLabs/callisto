package remote

import (
	"github.com/forbole/juno/v5/node/remote"
	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"

	storagesource "github.com/forbole/bdjuno/v4/modules/storage/source"
)

var _ storagesource.Source = &Source{}

// Source implements storagesource.Source using a remote node
type Source struct {
	*remote.Source
	querier storagetypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, querier storagetypes.QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetParams implements storagesource.Source
func (s Source) GetParams(height int64) (storagetypes.Params, error) {
	res, err := s.querier.Params(remote.GetHeightRequestContext(s.Ctx, height), &storagetypes.QueryParams{})
	if err != nil {
		return storagetypes.Params{}, nil
	}

	return res.Params, nil
}

// GetProviders implements storagesource.Source
func (s Source) GetProviders(height int64) ([]storagetypes.Providers, error) {
	res, err := s.querier.AllProviders(remote.GetHeightRequestContext(s.Ctx, height), &storagetypes.QueryAllProviders{})
	if err != nil {
		return []storagetypes.Providers{}, nil
	}

	return res.Providers, nil
}

// GetFiles implements storagesource.Source
func (s Source) GetFiles(height int64) ([]storagetypes.UnifiedFile, error) {
	res, err := s.querier.AllFiles(remote.GetHeightRequestContext(s.Ctx, height), &storagetypes.QueryAllFiles{})
	if err != nil {
		return []storagetypes.UnifiedFile{}, nil
	}

	return res.Files, nil
}

// ActiveProviders implements storagesource.Source
func (s Source) GetActiveProviders(height int64) ([]storagetypes.ActiveProviders, error) {
	res, err := s.querier.ActiveProviders(remote.GetHeightRequestContext(s.Ctx, height), &storagetypes.QueryActiveProviders{})
	if err != nil {
		return []storagetypes.ActiveProviders{}, nil
	}

	return res.Providers, nil
}

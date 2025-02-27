package types

import storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"

// StorageParams represents the x/storage parameters
type StorageParams struct {
	storagetypes.Params
	Height int64
}

// NewStorageParams allows to build a new StorageParams instance
func NewStorageParams(params storagetypes.Params, height int64) *StorageParams {
	return &StorageParams{
		Params: params,
		Height: height,
	}
}

// StorageProvider represents the x/storage providers
type StorageProvider struct {
	Address         string
	Ip              string
	Totalspace      string
	BurnedContracts string
	Creator         string
	KeybaseIdentity string
	AuthClaimers    []string
}

// NewStorageProvider allows to build a new StorageProvider instance
func NewStorageProvider(
	address string,
	ip string,
	totalspace string,
	burnedContracts string,
	creator string,
	keybaseIdentity string,
	authClaimers []string,
) *StorageProvider {
	return &StorageProvider{
		Address:         address,
		Ip:              ip,
		Totalspace:      totalspace,
		BurnedContracts: burnedContracts,
		Creator:         creator,
		KeybaseIdentity: keybaseIdentity,
		AuthClaimers:    authClaimers,
	}
}

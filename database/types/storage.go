package types

import storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"

// StorageParamsRow represents a single row of the "storage_params" table
type StorageParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}

// NewStorageParamsRow creates a new StorageParamsRow
func NewStorageParamsRow(
	params string, height int64,
) StorageParamsRow {
	return StorageParamsRow{
		OneRowID: true,
		Params:   params,
		Height:   height,
	}
}

// Equal reports whether m and n represent the same table rows.
func (m StorageParamsRow) Equal(n StorageParamsRow) bool {
	return m.Params == n.Params &&
		m.Height == n.Height
}

// NewStorageParams creates a new jackal x/storage Params instance
func NewStorageParams(depositAccount string, proofWindow int64, chunkSize int64,
	missesToBurn int64, priceFeed string, maxContractAgeInBlocks int64, pricePerTbPerMonth int64,
	attestFormSize int64, attestMinToPass int64, collateralPrice int64, checkWindow int64, polRatio int64, referralCommission int64,
) *storagetypes.Params {
	return &storagetypes.Params{
		DepositAccount:         depositAccount,
		ProofWindow:            proofWindow,
		ChunkSize:              chunkSize,
		MissesToBurn:           missesToBurn,
		PriceFeed:              priceFeed,
		MaxContractAgeInBlocks: maxContractAgeInBlocks,
		PricePerTbPerMonth:     pricePerTbPerMonth,
		AttestFormSize:         attestFormSize,
		AttestMinToPass:        attestMinToPass,
		CollateralPrice:        collateralPrice,
		CheckWindow:            checkWindow,
		PolRatio:               polRatio,
		ReferralCommission:     referralCommission,
	}
}

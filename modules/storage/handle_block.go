package storage

import (
	juno "github.com/forbole/juno/v5/types"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, res *tmctypes.ResultBlockResults, tx []*juno.Tx, vals *tmctypes.ResultValidators,
) error {
	// Update storage providers list
	go m.UpdateProviders(block.Block.Height)

	// Update active providers list
	go m.UpdateActiveProviders(block.Block.Height)

	return nil
}

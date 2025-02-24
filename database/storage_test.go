package database_test

import (
	"encoding/json"

	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"

	"github.com/forbole/bdjuno/v4/types"

	dbtypes "github.com/forbole/bdjuno/v4/database/types"
)

func (suite *DbTestSuite) TestBigDipperDb_SaveStorageParams() {
	storageParams := dbtypes.NewStorageParams(
		"jkl1DepositAccount",
		50,
		1024,
		3,
		"jklPrice",
		100,
		15,
		2,
		3,
		10_000_000_000,
		11,
		40,
		25,
	)
	err := suite.database.SaveStorageParams(types.NewStorageParams(*storageParams, 10))
	suite.Require().NoError(err)

	var rows []dbtypes.StorageParamsRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM storage_params`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)

	var storedParams *storagetypes.Params
	err = json.Unmarshal([]byte(rows[0].Params), &storedParams)
	suite.Require().NoError(err)
	suite.Require().Equal(storageParams, storedParams)
	suite.Require().Equal(int64(10), rows[0].Height)
}

package database_test

import (
	"encoding/json"

	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/lib/pq"

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

func (suite *DbTestSuite) TestBigDipperDb_SaveStorageProviders() {
	authClaimers := []string{
		"creator",
		"authClaimer1",
		"authClaimer2",
	}
	blockHeight := int64(10)

	storageProviderOne := storagetypes.Providers{
		Address:         "jkl1address12345678",
		Ip:              "198.162.1.3",
		Totalspace:      "1_000_000_000",
		BurnedContracts: "0",
		Creator:         "creator",
		KeybaseIdentity: "keybaseIdentity",
		AuthClaimers:    authClaimers,
	}
	storageProviderTwo := storagetypes.Providers{
		Address:         "jkl1address2222",
		Ip:              "198.162.1.3",
		Totalspace:      "1_000_000_000",
		BurnedContracts: "0",
		Creator:         "creator",
		KeybaseIdentity: "keybaseIdentity",
		AuthClaimers:    authClaimers,
	}

	storageProvidersList := []storagetypes.Providers{
		storageProviderOne,
		storageProviderTwo,
	}

	err := suite.database.SaveStorageProviders(storageProvidersList, blockHeight)
	suite.Require().NoError(err)

	var rows []dbtypes.StorageProviderRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM storage_providers`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 2)

	suite.Require().Equal(int64(10), rows[0].Height)
	suite.Require().Equal(pq.StringArray(authClaimers), rows[0].AuthClaimers)
}

func (suite *DbTestSuite) TestBigDipperDb_SaveActiveProviders() {
	activeProviders := []storagetypes.ActiveProviders{
		{Address: "jkl1ActiveProviderOne"},
		{Address: "jkl1ActiveProviderTwo"},
		{Address: "jkl1ActiveProviderThree"},
	}
	blockHeight := int64(10)

	err := suite.database.SaveActiveProviders(activeProviders, blockHeight)
	suite.Require().NoError((err))

	var rows []dbtypes.ActiveProviderRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM active_providers`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 3)

	for i, row := range rows {
		suite.Require().Equal(activeProviders[i].Address, row.Address)
	}
}

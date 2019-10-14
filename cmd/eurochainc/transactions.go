package main

import (
	eurochaintypes "github.com/nbh-digital/eurochain/pkg/types"
	"github.com/threefoldtech/rivine/extensions/authcointx"
	authcointxcli "github.com/threefoldtech/rivine/extensions/authcointx/client"
	"github.com/threefoldtech/rivine/extensions/minting"
	mintingcli "github.com/threefoldtech/rivine/extensions/minting/client"
	"github.com/threefoldtech/rivine/types"

	"github.com/threefoldtech/rivine/pkg/client"
)

func RegisterDevnetTransactions(bc *client.BaseClient) {
	registerTransactions(bc)
}

func RegisterTestnetTransactions(bc *client.BaseClient) {
	registerTransactions(bc)
}

func registerTransactions(bc *client.BaseClient) {
	// create minting plugin client...
	mintingCLI := mintingcli.NewPluginConsensusClient(bc)
	// ...and register minting types
	types.RegisterTransactionVersion(eurochaintypes.TransactionVersionMinterDefinition, minting.MinterDefinitionTransactionController{
		MintConditionGetter: mintingCLI,
		TransactionVersion:  eurochaintypes.TransactionVersionMinterDefinition,
	})
	types.RegisterTransactionVersion(eurochaintypes.TransactionVersionCoinCreation, minting.CoinCreationTransactionController{
		MintConditionGetter: mintingCLI,
		TransactionVersion:  eurochaintypes.TransactionVersionCoinCreation,
	})
	types.RegisterTransactionVersion(eurochaintypes.TransactionVersionCoinDestruction, minting.CoinDestructionTransactionController{
		TransactionVersion: eurochaintypes.TransactionVersionCoinDestruction,
	})

	// create coin auth tx plugin client...
	authCoinTxCLI := authcointxcli.NewPluginConsensusClient(bc)
	// ...and register coin auth tx types
	types.RegisterTransactionVersion(eurochaintypes.TransactionVersionAuthConditionUpdate, authcointx.AuthConditionUpdateTransactionController{
		AuthInfoGetter:     authCoinTxCLI,
		TransactionVersion: eurochaintypes.TransactionVersionAuthConditionUpdate,
	})
	types.RegisterTransactionVersion(eurochaintypes.TransactionVersionAuthAddressUpdate, authcointx.AuthAddressUpdateTransactionController{
		AuthInfoGetter:     authCoinTxCLI,
		TransactionVersion: eurochaintypes.TransactionVersionAuthAddressUpdate,
	})
}

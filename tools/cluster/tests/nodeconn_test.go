// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// TODO: Test connect/reconnect - start node conn, and later the hornet.
// TODO: Test connect/reconnect - on a running node stop and later restart hornet.

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/inx-app/nodebridge"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/l1connection"
	"github.com/iotaledger/wasp/packages/nodeconn"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/testutil/testlogger"
	"github.com/iotaledger/wasp/packages/testutil/testpeers"
	"github.com/iotaledger/wasp/packages/transaction"
)

func createChain(t *testing.T) *isc.ChainID {
	originator := cryptolib.NewKeyPair()
	layer1Client := l1connection.NewClient(l1.Config, testlogger.NewLogger(t))
	layer1Client.RequestFunds(originator.Address())
	utxoMap, err := layer1Client.OutputMap(originator.Address())
	require.NoError(t, err)

	var utxoIDs iotago.OutputIDs
	for id := range utxoMap {
		utxoIDs = append(utxoIDs, id)
	}

	originTx, chainID, err := transaction.NewChainOriginTransaction(
		originator,
		originator.Address(),
		originator.Address(),
		0,
		utxoMap,
		utxoIDs,
	)
	require.NoError(t, err)
	err = layer1Client.PostTx(originTx)
	require.NoError(t, err)

	return chainID
}

func TestNodeConn(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping nodeconn test in short mode")
	}

	l1.StartPrivtangleIfNecessary(t.Logf)

	log := testlogger.NewLogger(t)
	defer log.Sync()
	peerCount := 1

	//
	// Start a peering network.
	// peeringID := peering.RandomPeeringID()
	peerNetIDs, peerIdentities := testpeers.SetupKeys(uint16(peerCount))
	networkLog := testlogger.WithLevel(log.Named("Network"), 0, false)
	_, networkCloser := testpeers.SetupNet(
		peerNetIDs,
		peerIdentities,
		testutil.NewPeeringNetReliable(networkLog),
		networkLog,
	)
	t.Logf("Peering network created.")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	nodeBridge, err := nodebridge.NewNodeBridge(ctx, l1.Config.INXAddress, 10, log.Named("NodeBridge"))
	require.NoError(t, err)
	go nodeBridge.Run(ctx)

	nc := nodeconn.New(ctx, log, nodeBridge)

	//
	// Check the chain operations.
	chainID := createChain(t)
	chainOuts := make(map[iotago.OutputID]iotago.Output)
	chainOICh := make(chan iotago.OutputID)
	chainStateOuts := make(map[iotago.OutputID]iotago.Output)
	chainStateOutsICh := make(chan iotago.OutputID)
	mChan := make(chan *nodebridge.Milestone, 10)
	nc.RegisterChain(
		chainID,
		func(oi iotago.OutputID, o iotago.Output) {
			chainStateOuts[oi] = o
			chainStateOutsICh <- oi
		},
		func(oi iotago.OutputID, o iotago.Output) {
			chainOuts[oi] = o
			chainOICh <- oi
		},
		func(m *nodebridge.Milestone) {
			mChan <- m
		},
	)
	<-mChan

	client := l1connection.NewClient(l1.Config, log)
	// Post a TX directly, and wait for it in the message stream (e.g. a request).
	err = client.RequestFunds(chainID.AsAddress())
	require.NoError(t, err)
	t.Logf("Waiting for outputs posted via tangle...")
	oid := <-chainOICh
	t.Logf("Waiting for outputs posted via tangle... Done, have %v=%v", oid.ToHex(), chainOuts[oid])

	require.NoError(t, err)
	wallet := cryptolib.NewKeyPair()
	client.RequestFunds(wallet.Address())
	tx, err := l1connection.MakeSimpleValueTX(client, wallet, chainID.AsAddress(), 1*isc.Million)
	require.NoError(t, err)
	err = nc.PublishTransaction(chainID, tx)
	require.NoError(t, err)
	t.Logf("Waiting for outputs posted via nodeConn...")
	oid = <-chainOICh
	t.Logf("Waiting for outputs posted via nodeConn... Done, have %v=%v", oid.ToHex(), chainOuts[oid])

	nc.UnregisterChain(chainID)

	//
	// Cleanup.
	require.NoError(t, networkCloser.Close())
}

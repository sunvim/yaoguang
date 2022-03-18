/*
 * Copyright (C) 2022  mobus <sunsc0220@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package abci

import (
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/sunvim/yaoguang/blockchain"
	"github.com/sunvim/yaoguang/codes"
	"github.com/sunvim/yaoguang/txs"
	"github.com/sunvim/yaoguang/validators"
	"github.com/tendermint/tendermint/abci/types"
)

type Validators interface {
	validators.History
}

var _ types.Application = (*App)(nil)

type App struct {
	types.BaseApplication
	nodeInfo string
	// state
	blockchain    *blockchain.Blockchain
	validators    Validators
	mempoolLocker sync.Locker
	block         *types.RequestBeginBlock

	// fail gracefully
	panicFunc func(error)

	txsDecoder txs.Decoder
}

func NewApp(nodeInfo string, blockchain *blockchain.Blockchain, validators Validators, txsDecoder txs.Decoder) *App {
	app := &App{
		nodeInfo:   nodeInfo,
		blockchain: blockchain,
		validators: validators,
		txsDecoder: txsDecoder,
	}
	return app
}

// Info/Query Connection
// Return application info
func (app *App) Info(_ types.RequestInfo) types.ResponseInfo {
	return types.ResponseInfo{
		Data:             app.nodeInfo,
		Version:          "dev",
		LastBlockHeight:  int64(app.blockchain.LastBlockHeight()),
		LastBlockAppHash: app.blockchain.AppHashAfterLastBlock(),
	}
}

func (app *App) Query(reqQuery types.RequestQuery) types.ResponseQuery {
	defer func() {
		if r := recover(); r != nil {
			app.panicFunc(fmt.Errorf("panic occured in abci.App/Query"))
		}
	}()

	var rawResponse types.ResponseQuery
	rawResponse.Log = "Query not supported"
	rawResponse.Code = codes.UnsupportedRequestCode

	return rawResponse
}

// Mempool Connection
// Validate a tx for the mempool
func (app *App) CheckTx(req types.RequestCheckTx) (rsp types.ResponseCheckTx) {
	const logHeader = "CheckTx"
	defer func() {
		if r := recover(); r != nil {
			app.panicFunc(fmt.Errorf("panic occurred in abci.App/CheckTx: %v\n%s", r, debug.Stack()))
		}
	}()
	log.Info().Str("event", "entry").Msg(logHeader)
	log.Info().Str("event", "exit").Msg(logHeader)
	return
}

// Provide the Mempool lock. When provided we will attempt to acquire this lock in a goroutine during the Commit. We
// will keep the checker cache locked until we are able to acquire the mempool lock which signals the end of the commit
// and possible recheck on Tendermint's side.
func (app *App) SetMempoolLocker(mempoolLocker sync.Locker) {
	app.mempoolLocker = mempoolLocker
}

// Consensus Connection
// Initialize blockchain w validators/other info from TendermintCore
func (app *App) InitChain(req types.RequestInitChain) (rsp types.ResponseInitChain) {
	const logHeader = "InitChain"
	defer func() {
		if r := recover(); r != nil {
			app.panicFunc(fmt.Errorf("panic occurred in abci.App/InitChain: %v\n%s", r, debug.Stack()))
		}
	}()
	log.Info().Str("event", "entry").Msg(logHeader)
	log.Info().Str("event", "exit").Msg(logHeader)
	return
}

func (app *App) BeginBlock(req types.RequestBeginBlock) (rsp types.ResponseBeginBlock) {
	const logHeader = "BeginBlock"
	defer func() {
		if r := recover(); r != nil {
			app.panicFunc(fmt.Errorf("panic occurred in abci.App/BeginBlock: %v\n%s", r, debug.Stack()))
		}
	}()
	log.Info().Str("event", "entry").Msg(logHeader)
	log.Info().Str("event", "exit").Msg(logHeader)
	return
}

func (app *App) DeliverTx(req types.RequestDeliverTx) (rsp types.ResponseDeliverTx) {
	const logHeader = "DeliverTx"
	defer func() {
		if r := recover(); r != nil {
			app.panicFunc(fmt.Errorf("panic occurred in abci.App/DeliverTx: %v\n%s", r, debug.Stack()))
		}
	}()
	log.Info().Str("event", "entry").Msg(logHeader)
	log.Info().Str("event", "exit").Msg(logHeader)
	return
}

func (app *App) EndBlock(req types.RequestEndBlock) (rsp types.ResponseEndBlock) {
	const logHeader = "EndBlock"
	defer func() {
		if r := recover(); r != nil {
			app.panicFunc(fmt.Errorf("panic occurred in abci.App/EndBlock: %v\n%s", r, debug.Stack()))
		}
	}()
	log.Info().Str("event", "entry").Msg(logHeader)
	log.Info().Str("event", "exit").Msg(logHeader)
	return
}

func (app *App) Commit() (rsp types.ResponseCommit) {
	const logHeader = "Commit"
	defer func() {
		if r := recover(); r != nil {
			app.panicFunc(fmt.Errorf("panic occurred in abci.App/Commit: %v\n%s", r, debug.Stack()))
		}
	}()
	log.Info().Str("event", "entry").Msg(logHeader)
	log.Info().Str("event", "exit").Msg(logHeader)
	return
}

// State Sync Connection
// List available snapshots
func (app *App) ListSnapshots(_ types.RequestListSnapshots) types.ResponseListSnapshots {
	panic("not implemented") // TODO: Implement
}

func (app *App) OfferSnapshot(_ types.RequestOfferSnapshot) types.ResponseOfferSnapshot {
	panic("not implemented") // TODO: Implement
}

func (app *App) LoadSnapshotChunk(_ types.RequestLoadSnapshotChunk) types.ResponseLoadSnapshotChunk {
	panic("not implemented") // TODO: Implement
}

func (app *App) ApplySnapshotChunk(_ types.RequestApplySnapshotChunk) types.ResponseApplySnapshotChunk {
	panic("not implemented") // TODO: Implement
}

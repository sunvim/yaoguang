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

package blockchain

import (
	"sync"
	"time"

	"github.com/tendermint/tendermint/types"
	db "github.com/tendermint/tm-db"
)

type Blockchain struct {
	sync.RWMutex
	state              State
	db                 db.DB
	genesis            types.GenesisDoc
	lastBlockHash      []byte
	lastCommitTime     time.Time
	lastCommitDuration time.Duration
}

func NewBlockchain() *Blockchain {
	return &Blockchain{}
}

func (bc *Blockchain) SetGenesisDoc(genesis *types.GenesisDoc) {
	bc.genesis = *genesis
}

func (bc *Blockchain) LastBlockHeight() uint64 {

	if bc == nil {
		return 0
	}

	bc.RLock()
	defer bc.RUnlock()
	return bc.state.LastBlockHeight

}

func (bc *Blockchain) AppHashAfterLastBlock() []byte {
	bc.RLock()
	defer bc.RUnlock()
	return bc.state.AppHashAfterLastBlock
}

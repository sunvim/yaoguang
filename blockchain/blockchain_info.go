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
	"time"

	"github.com/tendermint/tendermint/types"
)

type BlockchainInfo interface {
	GenesisHash() []byte
	GenesisDoc() types.GenesisDoc
	ChainID() string
	LastBlockHeight() uint64
	LastBlockTime() time.Time
	LastCommitTime() time.Time
	LastCommitDuration() time.Duration
	LastBlockHash() []byte
	AppHashAfterLastBlock() []byte
	// BlockHash gets the hash at a height (or nil if no BlockStore mounted or block could not be found)
	BlockHash(height uint64) ([]byte, error)
	// GetBlockHeader returns the header at the specified height
	GetBlockHeader(blockNumber uint64) (*types.Header, error)
	// GetNumTxs returns the number of transactions included in a particular block
	GetNumTxs(blockNumber uint64) (int, error)
}

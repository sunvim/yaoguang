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

package core

import (
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/sunvim/yaoguang/abci"
	"github.com/sunvim/yaoguang/blockchain"
	abciclient "github.com/tendermint/tendermint/abci/client"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/libs/service"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/types"
)

type Kern struct {
	node service.Service
}

func NewKern(nodeInfo, config string) *Kern {
	kern := &Kern{}
	node, err := newTendermint(nodeInfo, config)
	if err != nil {
		panic(err)
	}
	kern.node = node
	return kern
}

func newTendermint(nodeInfo, configFile string) (service.Service, error) {
	var err error
	// read config
	config := cfg.DefaultConfig()
	config.RootDir = filepath.Dir(filepath.Dir(configFile))
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper failed to read config file")
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "viper failed to unmarshal config")
	}
	if err := config.ValidateBasic(); err != nil {
		return nil, errors.Wrap(err, "config is invalid")
	}

	genesisDoc, err := types.GenesisDocFromFile(config.GenesisFile())
	if err != nil {
		panic(err)
	}

	bc := blockchain.NewBlockchain()
	bc.SetGenesisDoc(genesisDoc)

	app := abci.NewApp(nodeInfo, bc, nil, nil)

	// create local client
	localClient := abciclient.NewLocalCreator(app)

	// create logger
	logger, err := log.NewDefaultLogger(config.LogFormat, config.LogLevel, false)

	// create node
	node, err := nm.New(
		config,
		logger, localClient, genesisDoc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new Tendermint node")
	}
	return node, nil
}

func (k *Kern) Boot() error {
	return k.node.Start()
}

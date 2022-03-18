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

package commands

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sunvim/utils/grace"
	"github.com/sunvim/yaoguang/core"
	"github.com/sunvim/yaoguang/share"
)

var CmdStart = &cobra.Command{
	Use:     "start",
	Aliases: []string{"node", "run"},
	Short:   "boot this service",
	Run: func(cmd *cobra.Command, args []string) {

		_, srv := grace.New(cmd.Context())

		log.Info().Str("service", "booting").Msg("start")
		config, _ := cmd.Flags().GetString(share.BootConfig)
		log.Info().Str("config", config).Msg("start")
		nodeInfo, _ := cmd.Flags().GetString(share.BootNodeInfo)
		kern := core.NewKern(nodeInfo, config)
		if err := kern.Boot(); err != nil {
			log.Fatal().Err(err).Msg("start")
		}

		srv.Wait()
	},
}

func init() {

	CmdStart.PersistentFlags().StringP(share.BootConfig, "c", "config/config.toml", "configuration for this service")
	CmdStart.PersistentFlags().StringP(share.BootNodeInfo, "", "dev", "node name or id")

	//bind flag
	viper.BindPFlag(share.BootConfig, CmdStart.Flags().Lookup(share.BootConfig))
	viper.BindPFlag(share.BootNodeInfo, CmdStart.Flags().Lookup(share.BootNodeInfo))

}

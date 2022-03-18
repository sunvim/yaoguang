/*
 * Copyright (C) 2022  mobus <sunsc0220@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
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
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Branch    = "main"
	Author    = "mobus"
	Email     = "<sunsc0220@gmail.com>"
	Date      = "2022-03-08"
	Commit    = "821288f"
	GoVersion = "go1.17.8 linux/amd64"
)

var CmdVersion = &cobra.Command{
	Use:   "version",
	Short: "show software verison",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("branch: ", Branch)
		fmt.Println("author: ", Author)
		fmt.Println("email: ", Email)
		fmt.Println("date: ", Date)
		fmt.Println("git commit: ", Commit)
		fmt.Println(GoVersion)
	},
}

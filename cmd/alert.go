/*
Copyright © 2024 Ameya Sathe <ameya.sathe@hotmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// alertCmd represents the alert command
var alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Run this command to setup alerting mechanism and pass params",
	Long: `This is not a standalone utility. It relies upon successful metric collection by the 'mcms metrics' command.
A future enhancement.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("alert called. Awaiting...")
	},
}

func init() {
	rootCmd.AddCommand(alertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

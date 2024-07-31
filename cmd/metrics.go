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

func init() {
	rootCmd.AddCommand(metricsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metricsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metricsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	metricsCmd.Flags().IntP("interval", "n", 10, "Specify update interval")

}

// metricsCmd represents the metrics command
var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Gather metrics",
	Long: `Gathers metrics at pre-defined interval and stores it in a database.
	Syntax:
	mcms metrics -n, --interval seconds
	For example:
	mcms metrics -n 20
	`,
	Run: collectMetrics,
}

func collectMetrics(cmd *cobra.Command, args []string) {
	interval, _ := cmd.Flags().GetInt("interval")

	// Print the interval
	fmt.Printf("Collecting metrics at interval of %d seconds\n", interval)

}

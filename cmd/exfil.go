package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exfilCmd represents the exfil command
var exfilCmd = &cobra.Command{
	Use:   "exfil",
	Short: "Run exfil like tests on current host",
	Long: `Run exfil like tests on current host`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exfil called")
	},
}

func init() {
	rootCmd.AddCommand(exfilCmd)

	exfilCmd.Flags().Bool("dns", false, "Runs DNS like exfil tests")
	exfilCmd.Flags().Bool("ping", false, "Runs ping like exfil tests")
}

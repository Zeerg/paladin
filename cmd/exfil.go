package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exfilCmd represents the exfil command
var exfilCmd = &cobra.Command{
	Use:   "exfil",
	Short: "Run exfil like tests on current server",
	Long: `Run exfil like tests on current server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exfil called")
	},
}

func init() {
	rootCmd.AddCommand(exfilCmd)

}

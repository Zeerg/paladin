package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infectionCmd represents the infection command
var infectionCmd = &cobra.Command{
	Use:   "infection",
	Short: "Start a local infection server and attempt to spread",
	Long: `Start a local infection server and attempt to spread`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("infection called")
	},
}

func init() {
	rootCmd.AddCommand(infectionCmd)

}

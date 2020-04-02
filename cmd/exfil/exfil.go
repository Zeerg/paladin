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
var exfilDNS = &cobra.Command{
	Use:   "dns",
	Short: "Run dns exfil like tests on current host",
	Long: `Run dns exfil like tests on current host`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dns exfil called")
	},
}

var exfilDNSClient = &cobra.Command{
	Use:   "client",
	Short: "Run exfil DNS test as a client",
	Long: `Run exfil DNS test as a client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dns exfil client called")
	},
}

var exfilDNSServer = &cobra.Command{
	Use:   "server",
	Short: "Run exfil DNS test as a server",
	Long: `Run exfil DNS test as a server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dns exfil server called")
	},
}

var exfilPing = &cobra.Command{
	Use:   "ping",
	Short: "Run ping exfil like tests on current host",
	Long: `Run ping exfil like tests on current host`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping exfil called")
	},
}

func init() {

	rootCmd.AddCommand(exfilCmd)
	exfilCmd.AddCommand(exfilDNS)
	exfilCmd.AddCommand(exfilPing)
	exfilDNS.AddCommand(exfilDNSClient)
	exfilDNS.AddCommand(exfilDNSServer)

}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

)
var (

	dhost string
	exfilFileName string
)
// exfilCmd represents the exfil command
var exfilCmd = &cobra.Command{
	Use:   "exfil",
	Short: "Run exfil like tests on current host",
	Long: `Run exfil like tests on current host`,
}
var exfilDNS = &cobra.Command{
	Use:   "dns",
	Short: "Run dns exfil like tests on current host",
	Long: `Run dns exfil like tests on current host`,
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
		pingExfil(dhost, exfilFileName)
	},
}

var exfilPingReceive = &cobra.Command{
	Use:   "receive",
	Short: "Packet capture ping requests and reassemble files",
	Long: `Packet capture ping requests and reassemble file`,
	Run: func(cmd *cobra.Command, args []string) {
		pingExfil(dhost, exfilFileName)
	},
}

func init() {
	//Root Command
	rootCmd.AddCommand(exfilCmd)

	//Sub Commands
	exfilCmd.AddCommand(exfilDNS)
	exfilCmd.AddCommand(exfilPing)
	exfilPing.AddCommand(exfilPingReceive)
	exfilDNS.AddCommand(exfilDNSClient)
	exfilDNS.AddCommand(exfilDNSServer)

	//Ping Exfil Flags
	exfilPing.Flags().StringVarP(&dhost, "destination", "d", "", "The Destination Host of the Ping")
	exfilPing.Flags().StringVarP(&exfilFileName, "file", "f", "", "The name of the file to send over ping")

}

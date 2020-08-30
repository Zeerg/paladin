package cmd

import (

	"github.com/spf13/cobra"

)
var (
	dhost string
	exfilFileName string
	runTime int32
	outFile string
	dnsPort int
	dnsHost string
	remoteDNSPort string
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
		dnsClientExfil(exfilFileName, dnsHost, remoteDNSPort)
	},
}

var exfilDNSServer = &cobra.Command{
	Use:   "server",
	Short: "Run exfil DNS test as a server",
	Long: `Run exfil DNS test as a server`,
	Run: func(cmd *cobra.Command, args []string) {
		runDNSServer(dnsPort)
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
		pingListen(outFile, ipListen, runTime)
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

	//Ping Reassemble flags
	exfilPingReceive.Flags().StringVarP(&ipListen, "ip", "i", "0.0.0.0", "The ip to listen on")
	exfilPingReceive.Flags().Int32VarP(&runTime, "runTime", "r", 1024, "How long to run the ping listener")
	exfilPingReceive.Flags().StringVarP(&outFile, "outfile", "o", "out.text", "The destination filename")

	//DNS Server Flags
	exfilDNSServer.Flags().IntVarP(&dnsPort, "dnsPort", "p", 5353, "The Port to lisen on")

	//DNS Client Flags
	exfilDNSClient.Flags().StringVarP(&dnsHost, "dnsHost", "n", "", "The Destination Host of the DNS Request")
	exfilDNSClient.Flags().StringVarP(&remoteDNSPort, "remoteDNSPort", "o", "", "The Destination Port of the DNS Request")
	exfilDNSClient.Flags().StringVarP(&exfilFileName, "file", "f", "", "The name of the file to send over DNS Exfil")

}

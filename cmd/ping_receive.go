package cmd

import (
	"fmt"
	"time"
	"encoding/hex"

    "github.com/google/gopacket"
	"github.com/google/gopacket/pcap"

)
var (
    captureDevice       string = "eth0"
    captureTime int32 = 1024
    promiscuous  bool   = false
    err          error
    timeout      time.Duration = 1 * time.Second
	handle       *pcap.Handle
	filter 		 string = "icmp"
)

// pingReassemble takes the payload and reassembles it. 
func pingReassemble(outFile, captureDevice string, captureTime int32) {

    // Open device
    handle, err = pcap.OpenLive(captureDevice, captureTime, promiscuous, timeout)
    check(err)
	defer handle.Close()
	
    // Set filter
    err = handle.SetBPFFilter(filter)
    check(err)
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
		appLayer := packet.ApplicationLayer();
		payload, err := hex.DecodeString(string(appLayer.Payload()))
		check(err)
        fmt.Println(string(payload))
    }

}
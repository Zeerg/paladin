package cmd

import (
	"fmt"
	"time"

    "github.com/google/gopacket"
	"github.com/google/gopacket/pcap"

)
var (
    captureDevice       string = "eth0"
    captureTime int32 = 1024
    promiscuous  bool   = false
    err          error
    timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
	filter 		 string = "icmp"
)

// pingReassemble
func pingReassemble(captureDevice string, captureTime int32) {

    // Open device
    handle, err = pcap.OpenLive(captureDevice, captureTime, promiscuous, timeout)
    check(err)
	defer handle.Close()
	
    // Set filter
    
    err = handle.SetBPFFilter(filter)
    check(err)
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        // Do something with a packet here.
        fmt.Println(packet)
    }

}
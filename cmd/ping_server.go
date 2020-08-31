package cmd

import (
    "golang.org/x/net/icmp"
    "reflect"
    "log"
    "bytes"
)

var (
    ipListen       string = "0.0.0.0"
    captureTime int32 = 1024
    promiscuous  bool   = false
    err          error
)

// pingListen waits for the ping at the address
func pingListen(outFile, ipListen string, captureTime int32) {
    // Listen for ping
    pkt, err := icmp.ListenPacket("ip4:1", ipListen)
    check(err)
    // Wait to get request
	for {
		buf := make([]byte, 1024)
		_, addr, _ := pkt.ReadFrom(buf)
		clientAddr := addr
		m, err := icmp.ParseMessage(1,buf)
        check(err)
        datBody := reflect.ValueOf(m.Body).Elem().FieldByName("Data").Bytes()
        b := bytes.Trim(datBody, "\x00")
        decodedText := hexDecode(b)
        log.Println(decodedText)
        log.Println(clientAddr)
	}
}

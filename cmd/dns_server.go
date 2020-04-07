package cmd

import (

	"net"

	"github.com/Zeerg/paladin/log"
	"golang.org/x/net/dns/dnsmessage"

)

var (

	ip string
	ok bool
	dnsRequest string
	m dnsmessage.Message
	query dnsmessage.Question

)

func runDNSServer(dnsPort int) {
	addr := net.UDPAddr{
		Port: dnsPort,
		IP:   net.ParseIP("0.0.0.0"),
	}
	u, _ := net.ListenUDP("udp", &addr)

	// Wait to get request on that port
	for {
		buf := make([]byte, 1024)
		_, addr, _ := u.ReadFromUDP(buf)
		clientAddr := addr
		err = m.Unpack(buf)
		check(err)
		query = m.Questions[0]
		log.Println(query)
		log.Println(clientAddr)
	}
}
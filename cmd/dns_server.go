package cmd

import (

	"net"

	//"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
	"github.com/Zeerg/paladin/log"
)
var (

	ip string
	ok bool
	dnsRequest string

)
func decodeDNSQuery(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {

	var dnsAnswer layers.DNSResourceRecord
	dnsAnswer.Type = layers.DNSTypeA
	dnsRequest = string(request.Questions[0].Name)
	log.Println(dnsRequest)
	
}
func runDNSServer(dnsPort int) {
	addr := net.UDPAddr{
		Port: dnsPort,
		IP:   net.ParseIP("0.0.0.0"),
	}
	u, _ := net.ListenUDP("udp", &addr)

	// Wait to get request on that port
	for {
		tmp := make([]byte, 1024)
		_, addr, _ := u.ReadFrom(tmp)
		clientAddr := addr
		log.Println(u)
		log.Println(clientAddr)
	}
}
package cmd
import (
	"fmt"
	"os"
)

// a Function that gets a host name/ip and filename and splits it then sends it over ping
// Ping has a min size of 64 bytes and max of MTU so 1500
// Gopacket has ping
func pingExfil(destination string, fileName string) {
	fmt.Printf("Going to ping")
	f, err := os.Open("/tmp/dat")
    check(err)
	fmt.Println(destination)
	fmt.Println(fileName)
}
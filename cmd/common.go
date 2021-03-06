package cmd

import (
    "io/ioutil"
    "encoding/hex"

    "github.com/Zeerg/paladin/log"

)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

// hexEncode takes a file and dumps it to bytes. It's not meant for large files
func hexEncode(fileName string) []byte {

	dat, err := ioutil.ReadFile(fileName)
	check(err)
	dstEnc := make([]byte, hex.EncodedLen(len(dat)))
	hex.Encode(dstEnc, dat)
	return dstEnc
}
func hexDecode(bytesObject []byte) string {
    dst := make([]byte, hex.DecodedLen(len(bytesObject)))
    n, err := hex.Decode(dst, bytesObject)
    check(err)
    encodedMessage := string(dst[:n])
    return encodedMessage
}

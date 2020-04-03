package cmd

import (
	"github.com/Zeerg/paladin/log"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}
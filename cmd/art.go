package cmd

import (
	"fmt"
	"log"
	"io/ioutil"
	"github.com/spf13/cobra"
	"github.com/rakyll/statik/fs"

	_ "github.com/Zeerg/paladin/statik"
)
var (
	atomic string
)

// artCmd represents the art command
var artCmd = &cobra.Command{
	Use:   "art",
	Short: "Run Atomic Red Team Attacks",
	Long: `Run attomic red team attacks to test security alerting`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Running atomic attack " + atomic)
		statikFS, err := fs.New()
		if err != nil {
			log.Fatal(err)
		}

		atomicDir := "/" + atomic + "/" + atomic + ".yaml"
		log.Printf("Opening " + atomicDir)
		r, err := statikFS.Open(atomicDir)
		if err != nil {
			log.Fatal("Atomic Not Found")
		}    
		defer r.Close()
		contents, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal("Failed to Read Atomic")
		}

		fmt.Println(string(contents))
	},
}

func init() {
	rootCmd.AddCommand(artCmd)

	artCmd.Flags().StringVarP(&atomic, "atomic", "a", "", "Atomic test to run")
}

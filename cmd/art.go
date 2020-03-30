package cmd

import (
	"fmt"
	"log"
	"io/ioutil"
	"github.com/spf13/cobra"
	"github.com/rakyll/statik/fs"

	_ "github.com/Zeerg/paladin/statik"
)

// artCmd represents the art command
var artCmd = &cobra.Command{
	Use:   "art",
	Short: "Run Atomic Red Team Attacks",
	Long: `Run attomic red team attacks to test security alerting`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Running atomic attack")
		statikFS, err := fs.New()
		if err != nil {
			log.Fatal(err)
		}
		
		// Access individual files by their paths.
		r, err := statikFS.Open("/T1003/T1003.yaml")
		if err != nil {
			log.Fatal(err)
		}    
		defer r.Close()
		contents, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(contents))
	},
}

func init() {
	rootCmd.AddCommand(artCmd)

	artCmd.Flags().StringP("atomic", "a", "YOUR TEST", "Specify the atomic to run")
}

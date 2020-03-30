package cmd

import (
	"fmt"
	"log"
	"io/ioutil"
	"github.com/spf13/cobra"
	"github.com/rakyll/statik/fs"
	"github.com/manifoldco/promptui"

	_ "github.com/Zeerg/paladin/statik"
)
var (
	atomic string
)

func yesNo() bool {
    prompt := promptui.Select{
        Label: "Select[Yes/No]",
        Items: []string{"Yes", "No"},
    }
    _, result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Prompt failed %v\n", err)
    }
    return result == "Yes"

}
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
		fmt.Println("Would You Like to Run This Attack?")
		attack := yesNo()
		if attack {
			log.Printf("Attacking")

		} else {
			log.Printf("Not Running Attack")
		}
	},
}

func init() {
	rootCmd.AddCommand(artCmd)

	artCmd.Flags().StringVarP(&atomic, "atomic", "a", "", "Atomic test to run")
}

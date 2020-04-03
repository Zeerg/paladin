package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/Zeerg/paladin/log"
	"github.com/spf13/cobra"
	"github.com/rakyll/statik/fs"
	"github.com/manifoldco/promptui"
	"github.com/fatih/color"

	//Blank import for statik filesystem
	_ "github.com/Zeerg/paladin/statik"

)

var (

	atomic string
	displayName string
	config map[string]interface{}
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
	Long: `Run Atomic red team attacks to test security alerting
tests can be found here https://github.com/redcanaryco/atomic-red-team/blob/master/atomics/index.md`,
	Run: func(cmd *cobra.Command, args []string) {

		if atomic == "" {
			log.Fatal("Please Specify Atomic Attach Technique")
		}

		log.Printf("Running atomic attack " + atomic)
		statikFS, err := fs.New()
		check(err)

		atomicDir := "/" + atomic + "/" + atomic + ".yaml"
		log.Printf("Opening " + atomicDir)
		r, err := statikFS.Open(atomicDir)
		check(err)
		defer r.Close()
		contents, err := ioutil.ReadAll(r)
		check(err)
    	err = yaml.Unmarshal(contents, &config)
    	check(err)
		
		displayName = config["display_name"].(string)
		atomicTests := config["atomic_tests"].([]interface{})
		
		for _, v := range atomicTests {
			command := v.(interface{}).(map[interface {}]interface{})["executor"].(map[interface {}]interface{})["command"].(string)
			name := v.(interface{}).(map[interface {}]interface{})["name"]
			log.Println("Would You Like to Run This Attack?")
			color.Green(name.(string))
			fmt.Println(command)
			attack := yesNo()
			if attack {
				commands := strings.Split(command, "\n")
				for _, attackString := range commands {
					cmd := exec.Command(attackString)
					cmd.Stdin = os.Stdin
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					check(err)
				}
			} else {
				log.Printf("Not Running")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(artCmd)

	artCmd.Flags().StringVarP(&atomic, "atomic", "a", "", "Atomic technique to run..ie T1003")
}

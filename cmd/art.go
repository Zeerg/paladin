package cmd

import (
	"log"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
	"github.com/rakyll/statik/fs"
	"github.com/manifoldco/promptui"

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
	Long: `Run Atomic red team attacks to test security alerting`,
	Run: func(cmd *cobra.Command, args []string) {

		if atomic == "" {
			log.Fatal("Please Specify Atomic Attach Technique")
		}

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

    	err = yaml.Unmarshal(contents, &config)
    	if err != nil {
        	log.Fatal("Failed to unmarshal YAML")
		}
		
		displayName = config["display_name"].(string)
		atomicTests := config["atomic_tests"].([]interface{})
		
		for _, v := range atomicTests {
			command := v.(interface{}).(map[interface {}]interface{})["executor"].(map[interface {}]interface{})["command"]
			log.Println("Would You Like to Run This Attack?")
			log.Println(command)

			attack := yesNo()
			if attack {
				log.Printf("Attacking\n")
			} else {
				log.Printf("Not Attacking")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(artCmd)

	artCmd.Flags().StringVarP(&atomic, "atomic", "a", "", "Atomic technique to run..ie T1003")
}

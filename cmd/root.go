package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "paladin",
	Short: "Run att&cks and test your security response.",
	Long: `

	Paladin
	https://github.com/Zeerg/paladin

           ▒▒            
        ▓▓▒▒▓▓░░        
▒▒▒▒▒▒▒▒▓▓▓▓██▓▓▓▓▓▓▓▓▓▓
▒▒▓▓▓▓▓▓▓▓▓▓██████████▓▓
▒▒▓▓▓▓▓▓▓▓▓▓██████████▓▓
▒▒▓▓▓▓▓▓▓▓▓▓██████████▓▓
▒▒▓▓▓▓▓▓▓▓▓▓██████████▓▓
▒▒▓▓▓▓▓▓▓▓▓▓████████▓▓░░
  ▒▒▓▓▓▓▓▓▓▓████████▓▓  
  ▒▒▓▓▓▓▓▓▓▓████████▓▓  
    ▒▒▓▓▓▓▓▓██████▓▓    
    ▓▓▒▒▓▓▓▓████▓▓░░    
        ▒▒▒▒██▓▓▓▓        
  
			
Run attacks against your systems using atomic red
team from Red Canary and other tools.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

}

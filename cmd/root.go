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
  
			
Run attacks against your systems using the att&ck framework.`,
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

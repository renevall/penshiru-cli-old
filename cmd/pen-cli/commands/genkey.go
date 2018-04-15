package commands

import (
	"fmt"

	"github.com/renevall/penshiru/crypto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genkey)
}

var genkey = &cobra.Command{
	Use:   "genkey",
	Short: "Print a random 256bits key",
	Long:  `Commands generate a random 256bits key and prints it to screen`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := crypto.GenerateRandomString(32)
		if err != nil {
			fmt.Println("Could not generate key")
		} else {
			fmt.Println("256bits key: ", key)
		}
	},
}

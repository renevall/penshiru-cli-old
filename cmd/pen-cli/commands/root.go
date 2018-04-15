package commands

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "pen-cli",
	Short: "Pen-cli provide util commands to deal with Penshiru App",
	Long: `Penshiru Cli allows you to execute commands for 
				  the day to day operations of Penshiru Site.
				  Complete documentation is available at http://penshiru.io/cli`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var (
	cfgFile     string
	projectBase string
	userLicense string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "github.com/renevall/penshiru", "base project directory")
	// rootCmd.PersistentFlags().StringP("author", "a", "René Vallecillo", "Author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	// rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		home = home + "/Development/Go/src/github.com/renevall/penshiru"

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".conf")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

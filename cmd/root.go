package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cygnus",
	Short: "Help you to secure your file directly in your repository",
	Long: `Cygnus is a CLI tools that will help you to encrypt/decrypt
your confidential files directly in you project repository.

Cygnus is come with access management, so you can use it with your
teams to manage your secret files (e.g. environment files, API keys, etc)
without any hassle.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.AddCommand(versionCmd)
    rootCmd.AddCommand(initCmd)
}



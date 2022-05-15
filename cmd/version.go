package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Cygnus",
  Long:  `All software has versions. This is Cygnus'`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Cygnus v0.1.0 -- HEAD")
  },
}

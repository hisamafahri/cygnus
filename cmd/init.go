package cmd

import (
	"fmt"
	"os"

	"github.com/hisamafahri/cygnus/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
  Use:   "init",
  Short: "Init cygnus in current repository",
  Run: func(cmd *cobra.Command, args []string) {
        isExist, err := utils.IsFolderExist(".cyg")
        if err != nil {
            fmt.Printf(" error: %s", err)
            return
        }

        if isExist {
            fmt.Println(" error: Cygnus already initialized in this repository")
            return
        }

        if err := os.Mkdir(".cyg/", os.ModePerm); err != nil {
            fmt.Printf(" error: %s", err)
            return
        }

        fmt.Println("This should be called only when cygnus successfully initialized")
  },
}

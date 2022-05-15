package cmd

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/hisamafahri/cygnus/model"
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

        initQs := "What is your app name?"
        groupName, err := utils.PromptText(&initQs)
        if err != nil {
            // failed to create/open the file
            fmt.Printf(" error: %s", err)
            return
        }

        if err := os.Mkdir(".cyg/", os.ModePerm); err != nil {
            fmt.Printf(" error: %s", err)
            return
        }
        // TODO: Create config file
        // TODO: Write the default value to the config file
        config := model.Config{}
        config.App.Name = groupName

        f, err := os.Create(".cyg/config.toml")
        if err != nil {
            // failed to create/open the file
            fmt.Printf(" error: %s", err)
            return
        }

        if err := toml.NewEncoder(f).Encode(config); err != nil {
            // failed to encode
            fmt.Printf(" error: %s", err)
            return
        }
        
        if err := f.Close(); err != nil {
            // failed to close the file
            fmt.Printf(" error: %s", err)
            return
        }
  },
}

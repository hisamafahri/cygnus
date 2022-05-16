package create

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/hisamafahri/cygnus/constant"
	"github.com/hisamafahri/cygnus/model"
	"github.com/hisamafahri/cygnus/utils"
)

func createGroup() {
    isExist, err := utils.IsFolderExist(".cyg")
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    if !isExist {
        fmt.Println(" error: Cygnus is not initialized in this repository")
        fmt.Println(" info: run 'cygnus init' to get started")
        return
    }

    var config model.Config

    _, err = toml.DecodeFile(constant.ConfigFilePath, &config)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    initQs := "What is your group name?"
    groupName, err := utils.PromptText(&initQs)
    if err != nil {
        // failed to create/open the file
        fmt.Printf(" error: %s", err)
        return
    }

    err = config.CreateGroup(&groupName)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    f, err := os.OpenFile(constant.ConfigFilePath, os.O_WRONLY, 0600)
    if err != nil {
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
}

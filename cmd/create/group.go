package create

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/hisamafahri/cygnus/constant"
	"github.com/hisamafahri/cygnus/model"
	"github.com/hisamafahri/cygnus/utils"
)

func createGroup() {
    err := utils.IsInitialized()
    if err != nil {
        fmt.Printf(" error: %s", err)
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
        fmt.Printf(" error: %s", err)
        return
    }

    err = config.CreateGroup(&groupName)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    err = utils.WriteToConfig(config)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }
}

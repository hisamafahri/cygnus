package create

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/hisamafahri/cygnus/constant"
	"github.com/hisamafahri/cygnus/model"
	"github.com/hisamafahri/cygnus/utils"
)

func createUser() {
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

    nameQs := "What is your name?"
    name, err := utils.PromptText(&nameQs)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    emailQs := "What is your email? (unique)"
    email, err := utils.PromptText(&emailQs)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    err = config.CreateUser(&name, &email)
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

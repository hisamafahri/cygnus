package create

import (
	"fmt"

	"github.com/hisamafahri/cygnus/utils"
)

func createUser() {
    err := utils.IsInitialized()
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    config, err := utils.DecodeConfig()

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

    err = utils.EncodeConfig(config)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }
}

package create

import (
	"fmt"

	"github.com/hisamafahri/cygnus/utils"
)

func createGroup() {
    err := utils.IsInitialized()
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    config, err := utils.DecodeConfig()

    groupQs := "What is your group name?"
    groupName, err := utils.PromptText(&groupQs)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    err = config.CreateGroup(&groupName)
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

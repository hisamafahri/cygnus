package add

import (
	"fmt"

	"github.com/hisamafahri/cygnus/utils"
)

func addFile() {
    err := utils.IsInitialized()
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    config, err := utils.DecodeConfig()

    // TODO: Autocomplete
    fileQs := "Which file you want to add?"
    filePath, err := utils.PromptText(&fileQs)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    // TODO: Ask as a multi select list
    groupQs := "Which group you want to add that file?"
    group, err := utils.PromptText(&groupQs)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }

    err = config.AddFile(filePath, &group)
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


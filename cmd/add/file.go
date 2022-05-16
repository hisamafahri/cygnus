package add

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/hisamafahri/cygnus/constant"
	"github.com/hisamafahri/cygnus/model"
	"github.com/hisamafahri/cygnus/utils"
)

func addFile() {
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

    err = utils.WriteToConfig(config)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return
    }
}

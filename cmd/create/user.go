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
}

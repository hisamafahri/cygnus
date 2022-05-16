package create

import (
    "fmt"
    "os"

    "github.com/BurntSushi/toml"
    "github.com/hisamafahri/cygnus/constant"
    "github.com/hisamafahri/cygnus/model"
    "github.com/hisamafahri/cygnus/utils"
    )

func createUser() {
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
}

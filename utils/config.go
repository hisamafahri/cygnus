package utils

import (
    "errors"
    "os"

    "github.com/BurntSushi/toml"
    "github.com/hisamafahri/cygnus/constant"
    "github.com/hisamafahri/cygnus/model"
    )

func WriteToConfig(config model.Config) error {
    f, err := os.OpenFile(constant.ConfigFilePath, os.O_WRONLY, 0600)
    if err != nil {
        return err
    }

    if err := toml.NewEncoder(f).Encode(config); err != nil {
        return err
    }

    if err := f.Close(); err != nil {
        return err
    }

    return nil
}

func IsInitialized() error {
    isExist, err := IsFolderExist(".cyg")
    if err != nil {
        return err
    }

    if !isExist {
        return errors.New("Cygnus is not initialize in this repository\n")
    }

    return nil
}

func WriteNewConfig(config model.Config) error {
    f, err := os.Create(constant.ConfigFilePath)
    if err != nil {
        // failed to create/open the file
        return err
    }

    if err := toml.NewEncoder(f).Encode(config); err != nil {
        return err
    }

    if err := f.Close(); err != nil {
        return err
    }

    return nil
}

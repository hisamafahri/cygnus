package model

import (
    "errors"
    )

type Config struct {
    App AppData
    Groups map[string]Group
}

type AppData struct {
    Name string
}

type Group struct {
    Files []string
    Users []string
}

func (data *Config) CreateGroup(name *string) (error) {
    var group map[string]Group

    if len(data.Groups) == 0 {
        group = make(map[string]Group)
    } else {
        group = data.Groups
    }

    if _, ok := group[*name]; ok {
        return errors.New("Group already exist\n")
    }
    group[*name] = Group {
        Files: []string{},
        Users: []string{},
    }
    data.Groups = group
    return nil
}

package model

import (
    "errors"
    )

type Config struct {
    App AppData
    Groups map[string]Group
    Users map[string]User
}

type AppData struct {
    Name string
}

type Group struct {
    Files []string
    Users []string
}

type User struct {
    Name string
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


func (data *Config) CreateUser(name, email *string) (error) {
    var user map[string]User

    if len(data.Users) == 0 {
        user = make(map[string]User)
    } else {
        user = data.Users
    }

    if _, ok := user[*email]; ok {
        return errors.New("User with that email already exist\n")
    }
    user[*email] = User {
        Name: *name,
    }
    data.Users = user
    return nil
}

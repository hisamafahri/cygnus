package model

import (
	"errors"
	"strings"
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
    Files []string // IMPORTANT: In lowercase
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

    if _, ok := user[strings.ToLower(*email)]; ok {
        return errors.New("User with that email already exist\n")
    }
    user[strings.ToLower(*email)] = User {
        Name: *name,
    }
    data.Users = user
    return nil
}

// TODO: GroupName input as a []string
func (data *Config) AddFile(filePath string, groupName *string) (error) {
    var files []string

    if len(data.Groups[*groupName].Files) == 0 {
        files = []string{}
    } else {
        files = data.Groups[*groupName].Files
    }

    // TODO: Loop through if GroupName input is []string
    files = append(files, filePath)

    data.Groups[*groupName] = Group {
        Files: files,
        Users: data.Groups[*groupName].Users,
    }
    return nil
}

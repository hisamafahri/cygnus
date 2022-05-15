package model

type Config struct {
    App AppData
    Groups map[string]Group
}

type AppData struct {
    Name string
}

type Group struct {
    files []string
    users []string
}

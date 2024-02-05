package main

import (
    toml "github.com/pelletier/go-toml/v2"
)


var doc string = `
IgnoreDir = [".git", "testDir"]
IgnoreFile = []
Editor = "nvim"
`

type Config struct {
    IgnoreDir  []string
    IgnoreFile []string
    Editor      string
}


func GetConfig() *Config {
    var cfg Config
    err := toml.Unmarshal([]byte(doc), &cfg)
    if err != nil {
        panic(err)
    }
    return &cfg
}










package config

import (
    "github.com/kelseyhightower/envconfig"
    "github.com/layeh/gumble/gumble"
)

var config Configuration

type Configuration struct {
    *gumble.Config
}

func configDefaults(c *Configuration) {
    c.Username = "mBot"
    c.Password = "mbot_password"
    c.Address = "localhost:64738"
}

func Get() *Config {
    c := &config
    configDefaults(c)
    envconfig.Process("mbot", c)
    return c
}

package config

import (
    "github.com/kelseyhightower/envconfig"
    "github.com/layeh/gumble/gumble"
    "os"
    "fmt"
    "github.com/davecgh/go-spew/spew"
)
//var Configuration gumble.Config

var Config Configuration

type Configuration struct {
    Insecure bool
    Username string // copy the values from our config to the embedded config
    Password string
    Address  string
    Gc gumble.Config //embed the Gumble config into our config
}

func gumbleConfigDefaults(c *Configuration) {
    c.Gc.Username = "mBot"
    c.Gc.Password = "mbot_password"
    c.Gc.Address = "localhost:64738"
}

func overrideDefaults(c *Configuration, gc *gumble.Config) {
    if c.Username != "" {
        c.Gc.Username = c.Username
    }
    if c.Password != "" {
        c.Gc.Password = c.Password
    }
    if c.Address != "" {
        c.Gc.Address = c.Address
    }
    if c.Insecure {
        c.Gc.TLSConfig.InsecureSkipVerify = true
    }
}

func init() {
    c := &Config
    /*
    var gumbleconfig gumble.Config
    */
    err := envconfig.Process("mbot", c)
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(2)
    }
    gumbleConfigDefaults(c)
    overrideDefaults(c, &c.Gc)
    spew.Dump(c)
}

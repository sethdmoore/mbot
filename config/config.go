package config

import (
    "github.com/kelseyhightower/envconfig"
    "github.com/layeh/gumble/gumble"
    "os"
    "fmt"
    //"github.com/davecgh/go-spew/spew"
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

func gumbleConfigDefaults(gc *gumble.Config) {
    gc.Username = "mBot"
    gc.Password = "mbot_password"
    gc.Address = "localhost:64738"
}

func overrideDefaults(c *Configuration, gc *gumble.Config) {
    if c.Username != "" {
        gc.Username = c.Username
    }
    if c.Password != "" {
        gc.Password = c.Password
    }
    if c.Address != "" {
        gc.Address = c.Address
    }
    if c.Insecure {
        gc.TLSConfig.InsecureSkipVerify = true
    }
}

func init() {
    c := &Config
    /*
    var gumbleconfig gumble.Config
    spew.Dump(gumbleconfig)
    */
    err := envconfig.Process("mbot", c)
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(2)
    }
    gumbleConfigDefaults(&c.Gc)
    overrideDefaults(c, &c.Gc)
}

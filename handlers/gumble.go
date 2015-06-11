package handlers

import (
    "fmt"
    "errors"
    "github.com/layeh/gumble/gumble"
    "github.com/sethdmoore/mbot/config"
    "github.com/davecgh/go-spew/spew"
)

func GumbleClient(config *config.Configuration) (c *gumble.Client, err error) {
    //var err error
    spew.Dump(config)
    fmt.Println("Hi")
    myclient := gumble.NewClient(&config.Gc)
    if myclient == nil {
        err = errors.New("Config for Gumble is empty. Could not create client")
        return nil, err
    }

    err = myclient.Connect()
    if err != nil {
        return nil, err
    }
    return myclient, err
}

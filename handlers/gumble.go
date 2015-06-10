package handlers

import (
    "errors"
    "github.com/layeh/gumble/gumble"
)

func GumbleClient() (c *gumble.Client, err error) {
    err error
    myclient := gumble.NewClient()
    if foo == nil {
        err = errors.New("Config for Gumble is empty. Could not create client")
        return err
    }

    err = gumble.Connect(myclient)
    if err != nil {
        return err
    }
    return myclient
}

package main

import (
    "fmt"
    "github.com/sethdmoore/mbot/handlers"
    "github.com/sethdmoore/mbot/config"
    "os"
    //"github.com/davecgh/go-spew/spew"
)

func main() {
    //spew.Dump(config.Config)
    this, err := handlers.GumbleClient(&config.Config)
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(2)
    }
    if this == nil {
        fmt.Printf("Oh no, it's nil")
        os.Exit(2)
    }
    fmt.Println("Hello")
    fmt.Printf("%+v", this)

}

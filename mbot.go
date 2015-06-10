package main

import (
    "fmt"
    "github.com/sethdmoore/mbot/handlers"
    "github.com/sethdmoore/mbot/config"
)

func main() {
    this := handlers.GumbleClient(&config)
    fmt.Println("Hello")

}

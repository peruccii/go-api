package main

import (
    "fmt"
    "log"
    
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    // todo: init routes passing router groups
    fmt.Println("main is running")
    
}

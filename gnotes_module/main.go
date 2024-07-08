package main

import (
    "flag"
    "fmt"
    "log"
    "os"

    h "test_module/handlers"

    "github.com/gofiber/fiber/v2"
)

const metadata = `
    {
        "name": "Gnotes", 
        "version": "1.0.0.0",
        "author": "0xUser",
        "data_type": 9
    }
`

func main() {
    port := flag.String("port", "", "The port on which the service will be launched")
    getInfo := flag.Bool("getinfo", false, "Checking if the module is running in the information retrieval mode")
    flag.Parse()

    if *getInfo == true {
        fmt.Print(metadata)
        os.Exit(0)
    }

    if *port == "" {
        log.Fatalf("Port not specified")
    }

    app := fiber.New()
    app.Post("/ExecuteModule", h.ExecuteModule)

    if err := app.Listen(":" + *port); err != nil {
        log.Fatalf("Error: %v", err)
    }
}

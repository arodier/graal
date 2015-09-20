package main

import (
    "log"
    "fmt"
    "net/http"
    "flag"
    "os"
    "os/signal"
)

// Manual import services and formatters, until plugins implemented
// Graal Specific services
import HelloService "./services/graal/hello/"

// System services entry point
import SystemService "./services/system/"

// Output formatters to use (json/xml/etc.)
import fmtJson "./formatters/"

const VERSION = "0.0.1"

// It is recemmended to use a port between 1180 and 1191 ;-)
var addressFlag = flag.String("ip", "127.0.0.1", "The IP address to bind. Use * for all")
var portFlag = flag.Int("port", 1188, "The port to listen on.")
var indentFlag = flag.Bool("indent", false, "Set to true if you want to indent the returned JSON")

func main() {

    // Initialise variables from the command line
    flag.Parse()

    var address = *addressFlag;
    var port = *portFlag;
    var indent = *indentFlag;

    // Start logging
    bind := fmt.Sprintf("http://%s:%d/", address, port)
    log.Printf("Starting the Graal server on '%s'", bind)

    // Initialise routes ================================

    // hello api
    http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
        // No parameters required for hello
        result := HelloService.Index("", nil)
        fmtJson.String(writer, request, result)
    })

    // home page
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        // Return nothing for now, we may add an option for user content later if requested
        writer.Header().Set("Content-Type", "text/html")
        writer.Write([]byte("Graal version "+VERSION))
    })

    // Load system services (monitor mem/load/etc.)
    SystemService.Load(indent)

    // Properly handle Ctrl-C
    channel := make(chan os.Signal, 1)
    signal.Notify(channel, os.Interrupt)
    go func() {
        for _ = range channel {
            log.Printf("Receiving interrupt. Exit")
            os.Exit(0)
        }
    }()

    // Start the server
    var error = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

    if error != nil {
        log.Fatal(error)
    }
}


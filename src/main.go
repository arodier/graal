package main

import (
    "log"
    "fmt"
    "net/http"
    "flag"
)

// Manual import services and formatters, until plugins implemented
// System services to import (input)
import HelloService "./services/graal/hello/"
import TimeService "./services/system/time/"
import StatsService "./services/system/stats/"

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
        // No parameters for hello
        result := HelloService.Index("", nil)
        fmtJson.String(writer, request, result)
    })

    // home page
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        // Return nothing for now, we may add an option for user content later if requested
        writer.Header().Set("Content-Type", "text/html")
        writer.Write([]byte("Graal version "+VERSION))
    })

    // system time functions
    http.HandleFunc("/system/time", func(writer http.ResponseWriter, request *http.Request) {
        // No parameters for hello
        result := TimeService.Index("", nil)
        fmtJson.Struct(writer, request, result, indent)
    })

    // system stats functions
    http.HandleFunc("/system/stats", func(writer http.ResponseWriter, request *http.Request) {
        // No parameters for hello
        result := StatsService.Index("", nil)
        fmtJson.Struct(writer, request, result, indent)
    })

    // Start the server
    var error = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

    if error != nil {
        log.Fatal(error)
    }
}

package main

import (
    "log"
    "fmt"
    "net/http"
    "flag"
    "os"
    "os/signal"
    "io/ioutil"
    "strings"
)

// Manual import services and formatters, until plugins implemented
import HelloService "./services/graal/hello/"    // Graal specific services
import SystemService "./services/system/"        // System services entry point
import NetworkService "./services/net/"          // Network services entry point

// Output formatters to use (json/xml/etc.)
import fmtJson "./formatters/"

const VERSION = "0.0.1"

// It is recemmended to use a port between 1180 and 1191 ;-)
var addressFlag = flag.String("ip", "127.0.0.1", "The IP address to bind. Use * for all")
var portFlag = flag.Int("port", 1188, "The port to listen on.")
var indentFlag = flag.Bool("indent", false, "Set to true if you want to indent the returned JSON")
var homeFlag = flag.String("home", "", "The root directory to answer other requests")

func main() {

    // Initialise variables from the command line
    flag.Parse()

    var address = *addressFlag;
    var port = *portFlag;
    var indent = *indentFlag;
    var home = *homeFlag;

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

    // Load system services (monitor mem/load/etc.)
    SystemService.Load(indent)

    // Load network services (monitor nic/hostname/firewall/etc.)
    NetworkService.Load(indent)

    // other static content
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

        if home != "" {

            home = strings.Trim(home, "/")
            path := request.URL.Path;
            if path == "" || path == "/" {
                path = "index.html"
            }

            // Serve content type
            ext := path[1+strings.LastIndex(path, "."):]
            content_type := ""
            switch ext {
            case "html":
                content_type = "text/html"
            case "js":
                content_type = "application/x-javascript"
            case "png":
                content_type = "image/png"
            case "jpeg", "jpg":
                content_type = "image/jpeg"
            default:
                // Not sure what to serve?
                content_type = "application/x-binary"
            }
            writer.Header().Set("Content-Type", content_type)

            fullPath := fmt.Sprintf("%s/%s", home, path)
            fullPath = strings.Replace(fullPath, "//", "/", -1)

            homeContent, error := ioutil.ReadFile(fullPath)

            if error == nil {
                writer.Write([]byte(homeContent))
                return
            } else {
                log.Printf("Cannot open home file from path '%s'", fullPath)
            }
        }

        // By default, just serve simple static content.
        writer.Write([]byte("Graal version "+VERSION))
    })

    // Properly handle Ctrl-C
    channel := make(chan os.Signal, 1)
    signal.Notify(channel, os.Interrupt)
    go func() {
        for _ = range channel {
            log.Printf("Receiving interrupt. Bye...")
            os.Exit(0)
        }
    }()

    // Start the server
    var error = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

    if error != nil {
        log.Fatal(error)
    }
}


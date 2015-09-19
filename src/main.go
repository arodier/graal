package main

import (
    "log"
    "fmt"
	"net/http"
)

// Manual import services and formatters, until plugins implemented
// System services to import (input)
import HelloService "./services/graal/hello/"

// Output formatters to use (json/xml/etc.)
import fmtJson "./formatters/"

// It is recemmended to use a port between 1180 and 1191 ;-)

func main() {

	// Initialise variables
	port := 1191
	address := "127.0.0.1"
	bind := fmt.Sprintf("http://%s:%d/", address, port)

	// Start logging
	log.Printf("Starting the Graal server on '%s'", bind)

	// Initialise routes
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		// No parameters for hello
		result := HelloService.Index("", nil)
		fmtJson.String(writer, request, result)
	})


	// Start the server
    var error = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if error != nil {
		log.Fatal(error)
	}
}

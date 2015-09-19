package json

import "encoding/json"
import "net/http"

// Exported symbols:
//export Simple

// Very simple string data format
func String(writer http.ResponseWriter, request *http.Request, data string) {

    // Default content type
    writer.Header().Set("Content-Type", "application/json")

	// A simple response type to encapsulate the data
    type SimpleResponse struct {
        Data      string
    }

	res := SimpleResponse {
		Data: data,
	}

	result, error := json.Marshal(res)

	if error == nil {
		// Default answer
		writer.Write([]byte(result))
	} else {
		writer.Write([]byte("Error"))
	}
}

// Format a more complex structure
func Struct(writer http.ResponseWriter, request *http.Request, data interface{}) {

    // Default content type
    writer.Header().Set("Content-Type", "application/json")

    // A simple response type to encapsulate the data
    type SimpleResponse struct {
        Data      interface {}
    }

    res := SimpleResponse {
        Data: data,
    }

    result, error := json.Marshal(res)

    if error == nil {
        // Default answer
        writer.Write([]byte(result))
    } else {
        writer.Write([]byte("Error"))
    }
}


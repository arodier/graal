
package hello

//export Index

import (
	"C"
)

func get() string {
	return "Hello, how are you?"
}


func Index(method string, params map[string]string) string {

	// Very simple example, ignore everything
	return get()
}


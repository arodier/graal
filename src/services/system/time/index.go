package time

//export Index

import (
    "C"
	"time"
)


func Index(method string, params map[string]string) time.Time {

    // Very simple, this is a read only method for now
	return time.Now()
}


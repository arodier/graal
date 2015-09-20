package system

import (
    "net/http"
)

// System services entry point
import TimeService "./time"
import StatsService "./stats"

// Output formatters to use (json/xml/etc.)
import fmtJson "../../formatters/"

func Load(indent bool) {

    // system time functions
    http.HandleFunc("/system/time", func(writer http.ResponseWriter, request *http.Request) {
        result := TimeService.Index("", nil)
        fmtJson.Struct(writer, request, result, indent)
    })

    // system stats functions
    http.HandleFunc("/system/stats", func(writer http.ResponseWriter, request *http.Request) {
        result := StatsService.Index("", nil)
        fmtJson.Struct(writer, request, result, indent)
    })

}

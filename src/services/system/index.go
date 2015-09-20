package system

import (
    "net/http"
)

// System services entry point
import TimeService     "./time"
import StatsService    "./stats"
import StorageService  "./storage"

// Output formatters to use (json/xml/etc.)
import fmtJson "../../formatters/"

func Load(indent bool) {

    // Local time on the server
    http.HandleFunc("/system/time", func(writer http.ResponseWriter, request *http.Request) {
        result := TimeService.Index("", nil)
        fmtJson.Struct(writer, request, result, indent)
    })

    // Statistics
    http.HandleFunc("/system/stats", func(writer http.ResponseWriter, request *http.Request) {
        result := StatsService.Index("", nil)
        fmtJson.Struct(writer, request, result, indent)
    })

    // Storage (disks/ssd)
    http.HandleFunc("/system/storage", func(writer http.ResponseWriter, request *http.Request) {
        result := StorageService.Index("", nil)
        fmtJson.Struct(writer, request, result, indent)
    })

}

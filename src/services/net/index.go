package net

import (
    "net/http"
)

// System services entry point
import BaseService     "./base"

// Output formatters to use (json/xml/etc.)
import fmtJson "../../formatters/"

func Load(indent bool) {

    // Basic methods
    http.HandleFunc("/net/hostname", func(writer http.ResponseWriter, request *http.Request) {
        result := BaseService.Hostname("", nil)
        fmtJson.Struct(writer, request, result, indent)
    })
}

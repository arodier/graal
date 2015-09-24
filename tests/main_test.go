package main_test

import (
    "testing"
    "net/http"
    "os"
)

func TestIndex(test *testing.T) {

    // TODO: Get the url from environment
    var url = os.Getenv("URL")
    if url == "" {
        url = "http://127.0.0.1:1188/"
    }

    resp, error := http.Get(url)

    if error != nil {
        test.Errorf("Cannot access home page: %v", error)
    } else {
        resp.Body.Close()
    }
}

package stats

//export Index

import (
    "C"
    "strings"
    "io/ioutil"
)

type systemStats struct {
    Load   []string
    Uptime []string
}

func Index(method string, params map[string]string) interface {} {

    stats := systemStats {}

    // Load average from /proc
    load, _ := ioutil.ReadFile("/proc/loadavg")
    stats.Load = strings.Split(strings.Replace(string(load), "\n", "", 1), " ")

    // uptime
    uptime, _ := ioutil.ReadFile("/proc/uptime")
    stats.Uptime = strings.Split(strings.Replace(string(uptime), "\n", "", 1), " ")

    return stats
}


package stats

//export Index

import (
    "C"
    "strings"
    "io/ioutil"
)

type MemInfo struct {
    Name string
    Value string
}

type systemStats struct {
    Load   []string
    Uptime []string
    MemInfo []MemInfo
}

func Index(method string, params map[string]string) interface {} {

    stats := systemStats {}

    // Load average from /proc
    // The first three fields in this file are load average figures giving
    // the number of jobs in the run queue (state R) or waiting for disk
    // I/O (state D) averaged over 1, 5, and 15 minutes.
    // They are  the  same  as  the  load average  numbers  given by uptime(1)
    // The fourth field consists of two numbers separated by a slash (/).
    // The first of these is the number of currently runnable kernel scheduling entities (processes, threads).
    // The second is the number of kernel scheduling entities that currently exist on the system.
    // The fifth field is the PID of the process that was most recently created on the system.
    load, error := ioutil.ReadFile("/proc/loadavg")
    if error == nil {
        stats.Load = strings.Split(strings.Replace(string(load), "\n", "", 1), " ")
    }

    // uptime: The first value is uptime, the second is idle time
    uptime, error := ioutil.ReadFile("/proc/uptime")
    if error == nil {
        stats.Uptime = strings.Split(strings.Replace(string(uptime), "\n", "", 1), " 1")
    }

    // uptime: The first value is uptime, the second is idle time
    memInfo, error := ioutil.ReadFile("/proc/meminfo")
    if error == nil {
        values := strings.Split(string(memInfo), "\n")
        size := len(values)
        stats.MemInfo = make([]MemInfo, size, size)
        for index,line := range values {
            nameValuePair := strings.Split(line, ":")
            if len(nameValuePair) == 2 && len(nameValuePair[0]) > 0 {
                stats.MemInfo[index].Name = nameValuePair[0]
                stats.MemInfo[index].Value = strings.Trim(nameValuePair[1], " ")
            }
        }
    }

    return stats
}


package stats

//export Index

import (
    "C"
    "strings"
    "strconv"
)

type LoadAvg struct {
    past1     float64
    past5     float64
    past15    float64
}

type Uptime struct {
    uptime    float64
    idle      float64
}

type MemBlockInfo struct {
    Name      string
    Value     float64
}

type SystemStats struct {
    Load     LoadAvg
    Uptime   Uptime
    MemInfo  []MemBlockInfo
}

func Index(method string, params map[string]string) SystemStats {

    var stats SystemStats

    // Load average from /proc
    // The first three fields in this file are load average figures giving
    // the number of jobs in the run queue (state R) or waiting for disk
    // I/O (state D) averaged over 1, 5, and 15 minutes.
    // They are  the  same  as  the  load average  numbers  given by uptime(1)
    // The fourth field consists of two numbers separated by a slash (/).
    // The first of these is the number of currently runnable kernel scheduling entities (processes, threads).
    // The second is the number of kernel scheduling entities that currently exist on the system.
    // The fifth field is the PID of the process that was most recently created on the system.
    load := GetLoadAvg()
    if load != "" {
        values := strings.Split(load, " ")
        stats.Load.past1, _ = strconv.ParseFloat(values[0], 1)
        stats.Load.past5, _ = strconv.ParseFloat(values[1], 1)
        stats.Load.past15, _ = strconv.ParseFloat(values[2], 1)
    }

    // uptime: The first value is uptime, the second is idle time
    uptime := GetUpTime()
    if uptime != "" {
        upAndIdle := strings.Split(uptime, " ")
        stats.Uptime.uptime, _ = strconv.ParseFloat(upAndIdle[0], 1)
        stats.Uptime.idle, _ = strconv.ParseFloat(upAndIdle[1], 1)
    }

    // uptime: The first value is uptime, the second is idle time
    memInfo := GetMemInfo()
    if memInfo != "" {
        values := strings.Split(string(memInfo), "\n")
        size := len(values)

        stats.MemInfo = make([]MemBlockInfo, size, size)

        for index,line := range values {

            if len(line) == 0 {
                continue
            }

            nameValuePair := strings.Split(line, ":")
            key := nameValuePair[0]
            value := strings.TrimSpace(strings.Replace(nameValuePair[1], "kB", "", 1))

            if len(key) > 0 && len(value) > 0 {
                stats.MemInfo[index].Name = key
                stats.MemInfo[index].Value, _ = strconv.ParseFloat(value, 1)
            }
        }
    }

    return stats
}


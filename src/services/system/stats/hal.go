// Very simple hardware abstraction layer
package stats

import (
    "io/ioutil"
    "strings"
)


func getProcEntry(path string) string {
    // TODO: probably log or something when an error occurs
    entry,error := ioutil.ReadFile(path)
    if error == nil {
        return strings.TrimSpace(string(entry))
    } else {
        return ""
    }
}

var GetLoadAvg = func() string {
    return getProcEntry("/proc/loadavg")
}

var GetUpTime = func() string {
    return getProcEntry("/proc/uptime")
}

var GetMemInfo = func() string {
    return getProcEntry("/proc/meminfo")
}


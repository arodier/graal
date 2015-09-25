package net

import (
    "strings"
    "os/exec"
    "bytes"
)

var GetHostName = func() string {

    var output bytes.Buffer
    command := exec.Command("/bin/hostname", "-f")
    command.Stdout = &output
    error := command.Run()

    if error == nil {
        return strings.TrimSpace(string(output.Bytes()))
    } else {
        return ""
    }
};

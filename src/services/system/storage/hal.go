package storage

import (
    "strings"
    "os/exec"
    "bytes"
)

var GetStorageInfo = func() string {

    var output bytes.Buffer
    command := exec.Command("/bin/df", "-l", "-T")
    command.Stdout = &output
    error := command.Run()

    if error == nil {
        return strings.TrimSpace(string(output.Bytes()))
    } else {
        return ""
    }
};

package time

import (
    "C"
    "strings"
    "regexp"
    "os/exec"
    "bytes"
    "log"
)

type DiskUsage struct {
    FileSystem    string
    Type          string
    Blocks        string
    Used          string
    Available     string
    PercentUsed   string
    Path          string
}

type StorageDetails struct {
    DiskUsage     []DiskUsage
}

func Index(method string, params map[string]string) interface {} {

    storageDetails := StorageDetails {}

    var output bytes.Buffer
    command := exec.Command("/bin/df", "-l", "-T")
    command.Stdout = &output
    error := command.Run()

    if error == nil {
        lines := strings.Split(string(output.Bytes()), "\n")
        size := len(lines)
        storageDetails.DiskUsage = make([]DiskUsage, size, size)

        for index,line := range lines {

            // Skip the headers and the final newline
            if index == 0 || index == size - 1 {
                continue
            }

            cells := regexp.MustCompile("\\s+").Split(line, 7)

            if len(cells) > 6 {
                storageDetails.DiskUsage[index-1].FileSystem = strings.Trim(cells[0], " ")
                storageDetails.DiskUsage[index-1].Type = strings.Trim(cells[1], " ")
                storageDetails.DiskUsage[index-1].Blocks = strings.Trim(cells[2], " ")
                storageDetails.DiskUsage[index-1].Used = strings.Trim(cells[3], " ")
                storageDetails.DiskUsage[index-1].Available = strings.Trim(cells[4], " ")
                storageDetails.DiskUsage[index-1].PercentUsed = strings.Trim(cells[5], " ")
                storageDetails.DiskUsage[index-1].Path = strings.Trim(cells[6], " ")
            }
        }
    } else {
        log.Fatal(error)
    }

    return storageDetails
}


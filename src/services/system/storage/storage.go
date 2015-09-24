package storage

import (
    "C"
    "strings"
    "regexp"
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

func Index(method string, params map[string]string) StorageDetails {

    storageDetails := StorageDetails {}

    storageInfo := GetStorageInfo()

    if storageInfo != "" {
        lines := strings.Split(storageInfo, "\n")
        size := len(lines) - 1
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
        log.Fatal("Cannot get storage info")
    }

    return storageDetails
}


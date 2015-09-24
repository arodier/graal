package storage

import (
    "testing"
    "fmt"
    "encoding/json"
    // "reflect"
    // "strconv"
)

var dfExample = `Filesystem              Type     1K-blocks      Used Available Use% Mounted on
/dev/dm-9               ext4       1886280   1241388    531024  71% /
udev                    devtmpfs     10240         0     10240   0% /dev
tmpfs                   tmpfs      3278456    353896   2924560  11% /run
/dev/dm-7               ext4      30626752  17927472  11120480  62% /usr
tmpfs                   tmpfs      8196132      3896   8192236   1% /dev/shm
tmpfs                   tmpfs         5120         4      5116   1% /run/lock
tmpfs                   tmpfs      8196132         0   8196132   0% /sys/fs/cgroup
/dev/mapper/system-var  ext4      61384724   7041072  51202432  13% /var`

func TestStorage(test *testing.T) {

    // Arrange
    // Mocking HAL
    GetStorageInfo = func() string { return dfExample }

    // Act
    var storageDetails StorageDetails = Index("", nil)

    // Asserts
    if len(storageDetails.DiskUsage) != 8 {
        test.Error("Wrong number of results: ")
        jsons, _ := json.MarshalIndent(storageDetails.DiskUsage, "", "  ")
        fmt.Printf("\n%s", jsons)
    }

    // TBC…
    // for var i=0 ; i<len(storageDetails) ; i++ {
    // }
}

package stats

import (
    "testing"
    "fmt"
    "encoding/json"
    "reflect"
    "strconv"
)

func TestIndex(test *testing.T) {

    // Arrange
    var laRef = LoadAvg {
        0.07,
        0.10,
        0.20,
    }

    var upRef = Uptime {
        2063420.79,
        14941349.92,
    }

    var memRef = [...]MemBlockInfo {
        { "MemTotal", 16392264 },
        { "MemFree", 260348 },
        { "MemAvailable", 7597448 },
        { "Buffers", 300860 },
        { "Cached", 7026568 },
        { "SwapCached", 15592 },
        { "Active", 10733472 },
    }

    // Mocking HAL
    GetLoadAvg = func() string { return "0.07 0.10 0.20 1/1000 474" }
    GetUpTime = func() string { return "2063420.79 14941349.92" }
    GetMemInfo = func() string {
        return "MemTotal:       16392264 kB\n" +
            "MemFree:          260348 kB\n" +
            "MemAvailable:    7597448 kB\n" +
            "Buffers:          300860 kB\n" +
            "Cached:          7026568 kB\n" +
            "SwapCached:        15592 kB\n" +
            "Active:         10733472 kB"
    }

    // Act
    var stats SystemStats = Index("", nil)

    // Asserts
    if stats.Load != laRef {
        test.Error("Wrong value for load")
    }

    if stats.Uptime != upRef {
        test.Error("Wrong value for uptime")
    }

    for i := 0 ; i < len(memRef) ; i++ {
        if !reflect.DeepEqual(stats.MemInfo[i], memRef[i]) {
            test.Error("Wrong value for memory on index " + strconv.Itoa(i))
            jsons, _ := json.MarshalIndent(stats.MemInfo[i], "", "  ")
            fmt.Printf("\n%s", jsons)
            jsons, _ = json.MarshalIndent(memRef[i], "", "  ")
            fmt.Printf("%s", jsons)
        }
    }
}

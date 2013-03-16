package reader

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func Read(name string) (G []int) {
    f, err := os.Open(name)
    if err != nil {
        return
    }
    r := bufio.NewReader(f)

    var n, m int
    var lineStr string
    var line []string
    for err == nil {
        lineStr, err = r.ReadString('\n')
        line = strings.Fields(lineStr)
        if len(line) < 1 || line[0] == "c" {
            continue
        }
        if line[0] == "p" {
            if len(line) != 4 {
                return
            }
            n, err = strconv.Atoi(line[2])
            m, err = strconv.Atoi(line[3])
        } else if line[0] == "a" {
            if len(line) != 4 {
                return
            }
            // make stuff
        }
    }
    fmt.Printf("%d %d\n",n, m)
    return
}

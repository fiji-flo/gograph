package main

import (
    "fmt"
    "bufio"
    "os"
)

func Read(name string) (G []int) {
    f, err := os.Open(name)
    if err != nil {
        return
    }
    r := bufio.NewReader(f)
    var line []byte 
    for err == nil {
        line, err = r.ReadBytes('\n')
        fmt.Printf("%s",line)
    }
    return
}

func main() {
    if len(os.Args) != 2 {
        return
    }
    Read(os.Args[1])
}



package main

import (
    "os"
    "github.com/fiji-flo/gograph/reader"
)

func main() {
    if len(os.Args) != 2 {
        return
    }
    reader.Read(os.Args[1])
}



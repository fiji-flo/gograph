package main 

import (
    "fmt"
    "os"
    "github.com/fiji-flo/gograph"
)

func main() {
    if len(os.Args) < 2 {
        return
    }
    G := gograph.Read(os.Args[1])
    fmt.Printf("%d %d\n", G.N, G.M)
    if len(os.Args) == 3 {
        gograph.Write(os.Args[2], G)
    }
}



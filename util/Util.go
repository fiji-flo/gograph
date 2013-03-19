package main 

import (
    "fmt"
    "os"
    "github.com/fiji-flo/gograph"
    "time"
)

func main() {
    if len(os.Args) < 2 {
        return
    }
    G := gograph.Read(os.Args[1])
    fmt.Printf("%d %d\n", G.N, G.M)
    var s, t int
    s = 1
    for {
        fmt.Printf("Enter s,t> ")
        fmt.Scanf("%d,%d", &s, &t)
        if s < 1 {
            break
        }
        tStart := time.Now()
        l := gograph.DijkstraST(&G, s, t)
        dur := time.Since(tStart)
        var µs = dur.Nanoseconds()/1000
        fmt.Printf("path: %d (in %dµs)\n", l, µs)
    }
    if len(os.Args) == 3 {
        gograph.Write(os.Args[2], G)
    }
}



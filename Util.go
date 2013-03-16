package gograph

import (
    "os"
)

func main() {
    if len(os.Args) != 2 {
        return
    }
    Read(os.Args[1])
}



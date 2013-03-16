package gograph

import (
    "fmt"
    "bufio"
    "os"
)

func Write(name string, G Graph) {
    f, err := os.Create(name)
    if err != nil {
        return
    }
    w := bufio.NewWriter(f)

    a := "a %d %d %d\n"

    w.WriteString(fmt.Sprintf("p cp %d %d\n", G.N, G.M))
    for i := 0; i < len(G.V); i++ {
        for _, v := range G.V[i].edges {
            w.WriteString(fmt.Sprintf(a, v.from, v.to, v.weight)) 
        }
    }
    f.Close()
}



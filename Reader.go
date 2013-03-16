package gograph

import (
    "bufio"
    "os"
    "strings"
    "strconv"
)

func Read(name string) (G Graph) {
    f, err := os.Open(name)
    if err != nil {
        return
    }
    r := bufio.NewReader(f)

    var V []Node
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
            V = make([]Node, n+1)
        } else if line[0] == "a" {
            if len(line) != 4 || V == nil {
                return
            }
            var e Edge
            e.from, err = strconv.Atoi(line[1])
            e.to, err = strconv.Atoi(line[2])
            e.weight, err = strconv.Atoi(line[3])

            V[e.from].id = e.from
            if V[e.from].edges == nil {
                V[e.from].edges = make(map[int]Edge)
            }
            V[e.from].edges[e.to] = e
            // make stuff
        }
    }
    f.Close()
    G.N = n
    G.M = m
    G.V = V
    return
}

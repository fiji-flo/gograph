package gograph

type Coords struct {
    x int
    y int
}

type Edge struct {
    from int
    to int
    weight int
    up bool
    short bool
    witness int
    orgEdges int
}

type Node struct {
    id int
    bits int
    co Coords
    edges map[int]Edge
    order int
}

type Graph struct {
    N int
    M int
    V []Node
}

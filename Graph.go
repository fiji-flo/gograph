package gograph

type Coords struct {
    x int
    y int
}

type Edge struct {
    from int
    to int
    weight int
}

type Node struct {
    id int
    bits int
    co Coords
    edges map[int]Edge
}

type Graph struct {
    N int
    M int
    V []Node
}

package gograph

type Edge struct {
    from int
    to int
    weight int
}

type Node struct {
    id int
    edges map[int]Edge
}

type Graph struct {
    N int
    M int
    V []Node
}

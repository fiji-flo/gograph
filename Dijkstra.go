package gograph

import (
	"container/heap"
)

func DijkstraST(g *Graph, sId int, tId int) int {
    s := &g.V[sId]

    const inf int = 1<<31 - 1

    pq := make(PriorityQueue, 0, g.N)

    for i := 1; i < g.N; i++ {
        g.V[i].bits = inf
    }
    s.bits = 0
    for vId, uv := range s.edges {
        item := &Item{
            value:      &g.V[vId],
            priority:   s.bits + uv.weight,
        }
        heap.Push(&pq, item)
    }
    cur := heap.Pop(&pq).(*Item)
    for cur.value.id != tId {
        u := cur.value
        if u.bits > cur.priority {
            u.bits = cur.priority
            for vId, uv := range u.edges {
                item := &Item{
                    value:      &g.V[vId],
                    priority:   u.bits + uv.weight,
                }
                heap.Push(&pq, item)
            }
        }
        cur = heap.Pop(&pq).(*Item)

    }
    return cur.priority
}




// An Item is something we manage in a priority queue.
type Item struct {
	value    *Node // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by changePriority and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	// To simplify indexing expressions in these methods, we save a copy of the
	// slice object. We could instead write (*pq)[i].
	a := *pq
	n := len(a)
	a = a[0 : n+1]
	item := x.(*Item)
	item.index = n
	a[n] = item
	*pq = a
}

func (pq *PriorityQueue) Pop() interface{} {
	a := *pq
	n := len(a)
	item := a[n-1]
	item.index = -1 // for safety
	*pq = a[0 : n-1]
	return item
}

// update is not used by the example but shows how to take the top item from
// the queue, update its priority and value, and put it back.
func (pq *PriorityQueue) update(value *Node, priority int) {
	item := heap.Pop(pq).(*Item)
	item.value = value
	item.priority = priority
	heap.Push(pq, item)
}

// changePriority is not used by the example but shows how to change the
// priority of an arbitrary item.
func (pq *PriorityQueue) changePriority(item *Item, priority int) {
	heap.Remove(pq, item.index)
	item.priority = priority
	heap.Push(pq, item)
}

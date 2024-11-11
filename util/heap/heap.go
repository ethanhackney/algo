package heap

// heap key value pair
type pair[K any, V any] struct {
    // key
    key   K
    // value
    value V
}

// binary heap
type Heap[K any, V any] struct {
    // heap of key value pairs
    arr []pair[K,V]
    // comparison function
    cmp func(a, b K) int
}

// create a new heap
func NewHeap[K any, V any](cmp func(a, b K) int) *Heap[K,V] {
    // return pointer to new heap
    return &Heap[K,V]{
        // heap of key value pairs
        arr: nil,
        // comparison function
        cmp: cmp,
    }
}

// push element onto heap
func (h *Heap[K,V]) Push(key K, value V) {
    // add new element to heap array
    h.arr = append(h.arr, pair[K,V]{key, value})
    // start heapify from new node
    i := len(h.arr) - 1
    // heapify up
    for i != 0 && h.cmp(h.arr[parent(i)].key, h.arr[i].key) > 0 {
        // swap parent and child
        swap(&h.arr[parent(i)], &h.arr[i])
        // move up tree
        i = parent(i)
    }
}

// get parent node
func parent(i int) int {
    // index of parent node
    return (i - 1) / 2
}

// swap values of two pointers
func swap[K any, V any](ip, jp *pair[K,V]) {
    // swap values
    *ip, *jp = *jp, *ip
}

// return top element of heap
func (h *Heap[K,V]) Top() (K, V) {
    // return key and value of root node
    return h.arr[0].key, h.arr[0].value
}

// pop top element off heap
func (h *Heap[K,V]) Pop() {
    // move last node to the root
    h.arr[0] = h.arr[len(h.arr) - 1]
    // shrink heap
    h.arr = h.arr[:len(h.arr) - 1]
    // start heapify down at root node
    i := 0
    // forever
    for {
        // minimum is root first
        min := i
        // get left child
        l := left(i)
        // get right child
        r := right(i)
        // if left child is less than root
        if l < len(h.arr) && h.cmp(h.arr[l].key, h.arr[i].key) < 0 {
            // left child becomes minimum
            min = l
        }
        // if right child is less than minimum
        if r < len(h.arr) && h.cmp(h.arr[r].key, h.arr[min].key) < 0 {
            // right child becomes minimum
            min = r
        }
        // if root is minimum
        if min == i {
            // we are done
            break
        }
        // swap root with minimum
        swap(&h.arr[i], &h.arr[min])
        // move down tree
        i = min
    }
}

// get left child
func left(i int) int {
    // index of left child
    return (i * 2) + 1
}

// get right child
func right(i int) int {
    // index of right child
    return (i * 2) + 2
}

// return whether heap is empty or not
func (h *Heap[K,V]) Empty() bool {
    // heap array has zero length
    return len(h.arr) == 0
}

// get size of heap
func (h *Heap[K,V]) Size() int {
    // return length of underlying array
    return len(h.arr)
}

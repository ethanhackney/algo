package main

import (
    "fmt"
    "util/heap"
)

// program to connect n ropes with minimal cost
func main() {
    // array of integers
    prices := []int{5, 4, 2, 8}
    // print result
    fmt.Printf("%d\n", mincost(prices))
}

// connect n ropes with minimal cost
func mincost(prices []int) int {
    // min heap
    pq := heap.NewHeap[int,bool](heapcmp)
    // add elements to heap
    for _, v := range prices {
        // push element onto heap
        pq.Push(v, true)
    }
    // minimum cost
    cost := 0
    // while heap has more than one element
    for pq.Size() > 1 {
        // extract top element
        x, _ := pq.Top()
        pq.Pop()
        // and new top element
        y, _ := pq.Top()
        pq.Pop()
        // get cost of extracted values
        sum := x + y
        // add new cost to heap
        pq.Push(sum, true)
        // update minimum cost
        cost += sum
    }
    // return minimal cost
    return cost
}

// comparison function for heap
func heapcmp(a, b int) int {
    return a - b
}

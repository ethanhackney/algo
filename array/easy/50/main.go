package main

import (
    "fmt"
    "util/heap"
)

// program to replace each array element by its corresponding rank
func main() {
    // array of integers
    nums := []int{10, 8, 15, 12, 6, 20, 1}
    // replace each element with rank
    rank(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// replace each array element by its corresponding rank
func rank(nums []int) {
    // max heap
    pq := heap.NewHeap[int,int](heapcmp)
    // add elements to heap
    for i, v := range nums {
        // add value and index to heap
        pq.Push(v, i)
    }
    // current rank
    rank := len(nums)
    // while heap non-empty
    for !pq.Empty() {
        // get top index
        _, i := pq.Top()
        // pop element off heap
        pq.Pop()
        // replace element with its rank
        nums[i] = rank
        // decrease rank
        rank--
    }
}

// reverse sort comparison function for max heap
func heapcmp(a, b int) int {
    // if greater, return -1
    if a > b {
        return -1
    }
    // if equal, return 0
    if a == b {
        return 0
    }
    // otherwise, return 1
    return 1
}

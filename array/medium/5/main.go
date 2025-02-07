package main

import "fmt"

// program to in-place merge two sorted arrays
func main() {
    // arrays of integers
    nums := [][]int{
        {1, 4, 7, 8, 10},
        {2, 3, 9},
    }
    // merge arrays
    merge(nums[0], nums[1])
    // print arrays
    fmt.Printf("%v, %v\n", nums[0], nums[1])
}

// in-place merge two sorted arrays
func merge(x, y []int) {
    // length of x
    m := len(x)
    // length of y
    n := len(y)
    // for each element of x
    for i := 0; i < m; i++ {
        // if current element is less than y[0]
        if x[i] < y[0] {
            // move ahead
            continue
        }
        // swap current element with first element of y
        swap(&x[i], &y[0])
        // new first element of y
        v := y[0]
        // move new element of y into correct position
        var k int
        for k = 1; k < n && y[k] < v; k++ {
            // move element down
            y[k - 1] = y[k]
        }
        // put new element into correct position
        y[k - 1] = v
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    // swap values
    *ip, *jp = *jp, *ip
}

package main

import "fmt"

// program to merge two arrays by satisfying given constraints
func main() {
    // arrays of integers
    nums := [][]int{
        {0, 2, 0, 3, 0, 5, 6, 0, 0},
        {1, 8, 9, 10, 15},
    }
    // merge arrays
    move(nums[0], nums[1])
    // print result
    fmt.Printf("%v\n", nums[0])
}

// move non-empty elements in x to beginning and merge with y
func move(x, y []int) {
    // next place to put non-empty element
    k := 0
    // for each element of x
    for _, v := range x {
        // if non-empty
        if v != 0 {
            // add to subarray
            x[k] = v
            // advance to next place
            k++
        }
    }
    // merge arrays
    merge(x, y, k - 1, len(y) - 1)
}

// merge x and y
func merge(x, y []int, m, n int) {
    // new size of x
    k := m + n + 1
    // while x and y have elements
    for m >= 0 && n >= 0 {
        // put next greater element in next free position
        if x[m] > y[n] {
            // add next element of x
            x[k] = x[m]
            // move to next element of x
            m--
        } else {
            // add next element of y
            x[k] = y[n]
            // move to next element of y
            n--
        }
        // move to next element of result
        k--
    }
    // copy remaining elements of y to x
    for n >= 0 {
        // add next element of y to x
        x[k] = y[n]
        // move to next element of x
        k--
        // move to next element of y
        n--
    }
}

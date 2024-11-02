package main

import "fmt"

// program to find duplicates within range k
func main() {
    // array of integers
    nums := []int{5, 6, 8, 2, 4, 6, 9}
    // range
    k := 4
    // if duplicates found
    if i := dup(nums, k); i != -1 {
        // print indices
        fmt.Printf("%d %d\n", i, i - k)
    }
}

// find duplicates within range k
func dup(nums []int, k int) int {
    // set to store elements within range k
    win := map[int]bool{}
    // for each element
    for i, v := range nums {
        // if v is in window, then duplicate found
        if _, ok := win[v]; ok {
            return i
        }
        // add current element to window
        win[v] = true
        // if i is greater than or equal to k
        if i >= k {
            // remove i - k element from window
            delete(win, i - k)
        }
    }
    // we reach here if no duplicates found
    return -1
}

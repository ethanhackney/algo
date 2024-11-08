package main

import "fmt"

// program to run binary search on an array
func main() {
    // array of integers
    nums := []int{2, 5, 6, 8, 9, 10}
    // number to search for
    n := 5
    // print result
    fmt.Printf("%d\n", bfind(nums, n))
}

// run binary search on array
func bfind(nums []int, n int) int {
    // low pointer of window
    low := 0
    // high pointer of window
    high := len(nums) - 1
    // while window is non-empty
    for low <= high {
        // get mid point
        mid := low + ((high - low) / 2)
        // if n found
        if n == nums[mid] {
            // return its index
            return mid
        } else if n < nums[mid] {
            // if n is less than mid, move high pointer down
            high = mid - 1
        } else {
            // if n is greater than mid, move low pointer up
            low = mid + 1
        }
    }
    // element not found
    return -1
}

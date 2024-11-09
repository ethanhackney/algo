package main

import "fmt"

// program to find first and last occurrence of element in a sorted array
func main() {
    // array of integers
    nums := []int{2, 5, 5, 5, 6, 6, 8, 9, 9, 9}
    // element to search for
    v := 5
    // print first occurrence
    fmt.Printf("%d\n", first(nums, v))
    // print last occurrence
    fmt.Printf("%d\n", last(nums, v))
}

// find first occurrence of element in sorted array
func first(nums []int, v int) int {
    // low pointer of window
    low := 0
    // high pointer of window
    high := len(nums) - 1
    // index of first occurrence
    fidx := -1
    // while window non-empty
    for low <= high {
        // get mid point
        mid := low + ((high - low) / 2)
        // if found v
        if v == nums[mid] {
            // update first occurrence to mid
            fidx = mid
            // move high pointer down
            high = mid - 1
        } else if v < nums[mid] {
            // if less than nums[mid], move high pointer down
            high = mid - 1
        } else {
            // if greater than nums[mid], move low pointer up
            low = mid + 1
        }
    }
    // return index of first occurrence if found
    return fidx
}

// find last occurrence of element in sorted array
func last(nums []int, v int) int {
    // low pointer of window
    low := 0
    // high pointer of window
    high := len(nums) - 1
    // index of last occurrence
    lidx := -1
    // while window non-empty
    for low <= high {
        // get mid point
        mid := low + ((high - low) / 2)
        // if v found
        if v == nums[mid] {
            // update last occurrence
            lidx = mid
            // move low pointer up
            low = mid + 1
        } else if v < nums[mid] {
            // if less than nums[mid], move high pointer down
            high = mid - 1
        } else {
            // if greater than nums[mid], move low pointer up
            low = mid + 1
        }
    }
    // return last occurrence if found
    return lidx
}

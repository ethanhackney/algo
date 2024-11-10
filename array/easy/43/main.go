package main

import "fmt"

// program to run exponential search on array
func main() {
    // array of integers
    nums := []int{2, 5, 6, 8, 9, 10}
    // value to search for
    v := 9
    // print result
    fmt.Printf("%d\n", efind(nums, v))
}

// run exponential search
func efind(nums []int, v int) int {
    // holds range in which v would be
    r := 1
    // find range in which v would be
    for r < len(nums) && nums[r] < v {
        // move to next power of 2
        r *= 2
    }
    // run binary searh on subarray [r/2, min(r, n - 1)]
    return bfind(nums, r / 2, min(r, len(nums) - 1), v)
}

// run binary search
func bfind(nums []int, low, high, v int) int {
    // while window non-empty
    for low <= high {
        // get mid point
        mid := low + ((high - low) / 2)
        // if v found
        if v == nums[mid] {
            // return mid
            return mid
        } else if v < nums[mid] {
            // if less than mid, then move high down
            high = mid - 1
        } else {
            // if greater than mid, then move low up
            low = mid + 1
        }
    }
    // v not found
    return -1
}

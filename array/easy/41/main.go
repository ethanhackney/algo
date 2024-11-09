package main

import "fmt"

// program to find the number of 1's in a sorted binary array
func main() {
    // sorted binary array
    nums := []int{0, 0, 0, 0, 1, 1, 1}
    // print result
    fmt.Printf("%d\n", count(nums))
}

// find number of 1's in a sorted binary array
func count(nums []int) int {
    // low pointer of window
    low := 0
    // high pointer of window
    high := len(nums) - 1
    // while window non-empty
    for low <= high {
        // get mid point
        mid := low + ((high - low) / 2)
        // if middle element is zero
        if nums[mid] == 0 {
            // move low pointer up
            low = mid + 1
        } else {
            // if end of zeros
            if mid == 0 || nums[mid - 1] == 0 {
                // return count of 1's
                return len(nums) - mid
            }
            // move high pointer down
            high = mid - 1
        }
    }
    // no 1's found
    return 0
}

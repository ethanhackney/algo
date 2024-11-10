package main

import "fmt"

// program to run interpolation search on array
func main() {
    // array of integers
    nums := []int{2, 5, 6, 8, 9, 10}
    // value to search for
    v := 5
    // print result
    fmt.Printf("%d\n", ifind(nums, v))
}

// run interpolation search
func ifind(nums []int, v int) int {
    // low pointer of window
    low := 0
    // high pointer of window
    high := len(nums) - 1
    // while in range that v can be found
    for nums[high] != nums[low] && v >= nums[low] && v <= nums[high] {
        // get estimated mid point
        mid := low + ((v - nums[low]) * (high - low) / (nums[high] - nums[low]))
        // if v found
        if v == nums[mid] {
            // return index
            return mid
        } else if v < nums[mid] {
            // if less than mid, move high pointer down
            high = mid - 1
        } else {
            // if greater than mid, move low pointer up
            low = mid + 1
        }
    }
    // if v found at low end
    if v == nums[low] {
        // return index
        return low
    }
    // v not found
    return -1
}

package main

import "fmt"

// program to find frequency of each element in sorted array with duplicates
func main() {
    // sorted array of integers
    nums := []int{2, 2, 2, 4, 4, 4, 5, 5, 6, 8, 8, 9}
    // print result
    fmt.Printf("%v\n", freq(nums))
}

// find frequency of each element in sorted array
func freq(nums []int) map[int]int {
    // frequency of elements
    f := map[int]int{}
    // low pointer of window
    low := 0
    // high pointer of window
    high := len(nums) - 1
    // while window non-empty
    for low <= high {
        // if window consists of one element
        if nums[low] == nums[high] {
            // update frequency
            f[nums[low]] += high - low + 1
            // move low pointer up
            low = high + 1
            // move high pointer up
            high = len(nums) - 1
        } else {
            // shrink window
            high = low + ((high - low) / 2)
        }
    }
    // return frequencies
    return f
}

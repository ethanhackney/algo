package main

import (
    "fmt"
    "math"
)

// program to find pair of elements with minimum absolute sum
func main() {
    // sorted array of integers
    nums := []int{-6, -5, -3, 0, 2, 4, 9}
    // find pair
    i, j := pair(nums)
    // print result
    fmt.Printf("%d, %d\n", nums[i], nums[j])
}

// find pair of elements with minimum absolute sum
func pair(nums []int) (int, int) {
    // low pointer
    low := 0
    // high pointer
    high := len(nums) - 1
    // minimum absolute sum
    minsum := math.MaxInt
    // indices of elements with minimum absolute sum
    idx := [2]int{}
    // while search space is not exhausted
    for low < high {
        // if absolute sum is less than minsum
        if abs(nums[low] + nums[high]) < minsum {
            minsum = abs(nums[low] + nums[high])
            idx[0] = low
            idx[1] = high
        }
        // increment low if less than 0
        if (nums[low] + nums[high]) < 0 {
            low++
        } else {
            // decrement high if more than 0
            high--

        }
    }
    // return pair
    return idx[0], idx[1]
}

// return absolute value of number
func abs(n int) int {
    // if less than zero
    if n < 0 {
        // negate it
        n = -n
    }
    // return result
    return n
}

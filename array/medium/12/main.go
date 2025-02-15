package main

import (
    "fmt"
    "math"
)

// program to find max difference of 2 array elements that satisfy constraint
func main() {
    // array of integers
    nums := []int{2, 7, 9, 5, 1, 3, 5}
    // get indices of elements with max difference
    i, j := diff(nums)
    // print elements
    fmt.Printf("%d, %d\n", nums[i], nums[j])
}

// find max difference of 2 array elements that satisfy constraint
func diff(nums []int) (int, int) {
    // indices of elements with max difference
    idx := [2]int{}
    // maximum difference
    maxdiff := math.MinInt
    // index of maximum element
    curmax := len(nums) - 1
    // from second to last element to first
    for i := len(nums) - 2; i >= 0; i-- {
        // if new maximum element
        if nums[i] >= nums[curmax] {
            // update maximum element
            curmax = i
        } else {
            // if new maximum difference
            if maxdiff < nums[curmax] - nums[i] {
                // update maximum difference
                maxdiff = nums[curmax] - nums[i]
                // update smaller element's index
                idx[0] = i
                // and greater element
                idx[1] = curmax
            }
        }
    }
    // return indices of pair with maximum difference
    return idx[0], idx[1]
}

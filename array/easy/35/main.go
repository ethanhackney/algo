package main

import (
    "fmt"
    "math"
)

// program to find elements in an array that are
// greater than all elements to the right
func main() {
    // array of integers
    nums := []int{10, 4, 6, 3, 5}
    // print result
    fmt.Printf("%v\n", gtr(nums))
}

// find elements that are greater than all elements to the right
func gtr(nums []int) []int {
    // holds current maximum element
    max := math.MinInt
    // holds indices of elements that are greater than all to the right
    idx := []int{}
    // for each element in reverse
    for i := len(nums) - 1; i >= 0; i-- {
        // if current element greater than maximum
        if nums[i] > max {
            // update maximum element
            max = nums[i]
            // add element to indices list
            idx = append(idx, i)
        }
    }
    // return indices
    return idx
}

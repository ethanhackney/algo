package main

import (
    "fmt"
    "math"
)

// program to find maximum product of two integers in an array
func main() {
    // array of integers
    nums := []int{-10, -3, 5, 6, -2}
    // get indices of values in maximum product
    i, j := maxprod(nums)
    // print result
    fmt.Printf("%d, %d\n", nums[i], nums[j])
}

// find maximum product of two integers in an array
func maxprod(nums []int) (int, int) {
    // indices of largest and second largest value
    maxidx := [2]int{0, -1}
    // indices of smallest and second smallest value
    minidx := [2]int{0, -1}
    // values of largest and second largest value
    maxval := [2]int{nums[0], math.MinInt}
    // values of smallest and second smallest value
    minval := [2]int{nums[0], math.MaxInt}
    // for each element other than first
    for i := 1; i < len(nums); i++ {
        // if larger than largest
        if nums[i] > maxval[0] {
            // set maxidx[1] to index of last largest value
            maxidx[1] = maxidx[0]
            // and maxidx[0] to index of new largest
            maxidx[0] = i
            // set maxval[1] to last largest value
            maxval[1] = maxval[0]
            // and maxval[0] to new largest value
            maxval[0] = nums[i]
        } else if nums[i] > maxval[1] {
            // update second largest index
            maxidx[1] = i
            // update second largest value
            maxval[1] = nums[i]
        }
        // if smallest than smallest
        if nums[i] < minval[0] {
            // set minidx[1] to index of last smallest value
            minidx[1] = minidx[0]
            // and minidx[2] to index of new smallest value
            minidx[0] = i
            // set minval[1] to last smallest value
            minval[1] = minval[0]
            // and minval[0] to new smallest value
            minval[0] = nums[i]
        } else if nums[i] < minval[1] {
            // update second smallest index
            minidx[1] = i
            // update second smallest value
            minval[1] = nums[i]
        }
    }
    // return maximum product
    if maxval[0] * maxval[1] > minval[0] * minval[1] {
        // return positive * positive indices
        return maxidx[0], maxidx[1]
    }
    // return negative * negative indices
    return minidx[0], minidx[1]
}

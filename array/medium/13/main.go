package main

import (
    "fmt"
    "math"
)

// program to print subarray with maximum sum
func main() {
    // array of integers
    nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    // get maximum sum subarray
    i, j := maxsum(nums)
    // print subarray
    fmt.Printf("%v\n", nums[i:j])
}

func maxsum(nums []int) (int, int) {
    // maximum sum
    maxsum := math.MinInt
    // maximum sum ending at current position
    maxhere := 0
    // indices of maximum sum subarray
    idx := [2]int{0, 0}
    // starting index of positive-sum sequence
    pos := 0
    // for each element
    for i, v := range nums {
        // update maximum sum ending here
        maxhere += v
        // if maxhere is less than current element
        if maxhere < v {
            // start from current element
            maxhere = v
            // update positive-sum starting index
            pos = i
        }
        // if new maximum sum subarray found
        if maxsum < maxhere {
            // update maximum
            maxsum = maxhere
            // update starting index
            idx[0] = pos
            // update ending index
            idx[1] = i
        }
    }
    // return subarray indices
    return idx[0], idx[1] + 1
}

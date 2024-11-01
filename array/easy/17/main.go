package main

import "fmt"

// program to find indices that divide array into two subarrays with equal sum
func main() {
    // array of integers
    nums := []int{-1, 6, 3, 1, -2, 3, 3}
    // print result
    for _, v := range divide(nums) {
        // print dividing index
        fmt.Printf("%v\n", v)
    }
}

// find indices that divides array into two subarrays with equal sum
func divide(nums []int) []int {
    // sum of all elements
    allsum := 0
    // calculate sum of all array elements
    for _, v := range nums {
        // add element to sum
        allsum += v
    }
    // indices that divide array into equal sum
    idx := []int{}
    // store sum of current subarray
    cursum := nums[0]
    // for each element after first
    for i := 1; i < len(nums) - 1; i++ {
        // if sum of nums[0...i-1] equals nums[i+1,n-1]
        if cursum == allsum - (nums[i] + cursum) {
            // add dividing index
            idx = append(idx, i)
        }
        // add element to current sum
        cursum += nums[i]
    }
    // return indices
    return idx
}

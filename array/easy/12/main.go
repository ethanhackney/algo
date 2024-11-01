package main

import (
    "fmt"
    "sort"
)

// program to find maximum sum formed by array digits
func main() {
    // array of integers
    nums := []int{4, 6, 2, 7, 9, 8}
    // find maximum sum formed by array digits
    odd, even := maxsum(nums)
    // print result
    fmt.Printf("%d, %d\n", odd, even)
}

// find maximum sum formed by array digits
func maxsum(nums []int) (int, int) {
    // sort array in descending order
    sort.Slice(nums, func(i, j int) bool {
        // reverse sort
        return nums[i] > nums[j]
    })
    // fill odd with digits at odd indices
    odd := 0
    // for every other element starting at 0
    for i := 0; i < len(nums); i += 2 {
        // add nums[i] to odd
        odd = odd * 10 + nums[i]
    }
    // fill even with digits at even indices
    even := 0
    // for every other element starting at 1
    for i := 1; i < len(nums); i += 2 {
        // add nums[i] to even
        even = even * 10 + nums[i]
    }
    // return odd and even
    return odd, even
}

package main

import "fmt"

// program to check if a subarray with 0 sum exists or not
func main() {
    // array of integers
    nums := []int{4, 2, -3, -1, 0, 4}
    // print result
    fmt.Printf("%v\n", zerosum(nums))
}

// check if a subarray with 0 sum exists or not
func zerosum(nums []int) bool {
    // stores sum of each subarray
    subsum := map[int]bool{0: true}
    // running sum
    sum := 0
    // for each element
    for _, v := range nums {
        // add element to sum
        if _, ok := subsum[v]; ok {
            // found zero sum
            return true
        }
        // add sum to set
        subsum[sum] = true
    }
    // no zero sum found
    return false
}

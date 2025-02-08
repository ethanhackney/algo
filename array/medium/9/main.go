package main

import "fmt"

// program to find largest subarray formed by consecutive integers
func main() {
    // array of integers
    nums := []int{2, 0, 2, 1, 4, 3, 1, 0}
    // find largest subarray
    i, j := maxsub(nums)
    // print subarray
    fmt.Printf("%v\n", nums[i:j])
}

// find largest subarray formed by consecutive integers
func maxsub(nums []int) (int, int) {
    // max length subarray
    maxlen := 1
    // index of start of longest subarray
    start := 0
    // index of end of longest subarray
    end := 0
    // for each subarray
    for i := 0; i < len(nums) - 1; i++ {
        // minimum value in subarray
        minval := nums[i]
        // maximum value in subarray
        maxval := nums[i]
        // for each element in this subarray
        for j := i + 1; j < len(nums); j++ {
            // update minimum value
            minval = min(minval, nums[j])
            // update maximum value
            maxval = max(maxval, nums[j])
            // check if subarray is consecutive
            if consecutive(nums, i, j, minval, maxval) {
                // if longest subarray
                if maxlen < maxval - minval + 1 {
                    // update max length
                    maxlen = maxval - minval + 1
                    // update start of subarray
                    start = i
                    // update end of subarray
                    end = j
                }
            }
        }
    }
    // return subarray indices
    return start, end + 1
}

// check if subarray is consecutive
func consecutive(nums []int, i, j, min, max int) bool {
    // is not consecutive if (max - min) is not equal to subarray length
    if (max - min) != (j - i) {
        // not consecutive
        return false
    }
    // keeps track of seen elements
    seen := make([]bool, j - i + 1)
    // check if each element appears once
    for k := i; k <= j; k++ {
        // if element seen
        if seen[nums[k] - min] {
            // not consecutive
            return false
        }
        // mark element as seen
        seen[nums[k] - min] = true
    }
    // all elements distinct
    return true
}

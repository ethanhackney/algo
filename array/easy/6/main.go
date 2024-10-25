package main

import "fmt"

// program to find maximum sum subarray
func main() {
    // array of integers
    nums := []int{2, 3, -8, 7, -1, 2, 3}
    // get maximum sum subarray indices
    i, j := maxsub(nums)
    // print subarray
    fmt.Printf("%v\n", nums[i:j+1])
}

// find maximum sum subarray
func maxsub(nums []int) (int, int) {
    // stores maximum sum found so far
    maxsum := nums[0]
    // stores maximum sum ending at current position
    maxhere := nums[0]
    // indices of first and last element in max sum subarray
    idx := [2]int{0, 0}
    // base of current subarray
    base := 0
    // for each element past first
    for i := 1; i < len(nums); i++ {
        // if starting a new subarray has greater sum than extending previous
        if maxhere + nums[i] < nums[i] {
            // update maxhere
            maxhere = nums[i]
            // reset base index to i
            base = i
        } else {
            // otherwise, add element to sum
            maxhere += nums[i]
        }
        // if the current subarray sum is greater than maxsum
        if maxhere > maxsum {
            // update maximum
            maxsum = maxhere
            // update start and end indices
            idx[0] = base
            idx[1] = i
        }
    }
    // return maximum sum subarray indices
    return idx[0], idx[1]
}

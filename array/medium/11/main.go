package main

import "fmt"

// program to find longest bitonic subarray
func main() {
    // array of integers
    nums := []int{3, 5, 8, 4, 5, 9, 10, 8, 5, 3, 4}
    // get longest bitonic subarray indices
    i, j := bitonic(nums)
    // print subarray
    fmt.Printf("%v\n", nums[i:j])
}

func bitonic(nums []int) (int, int) {
    // ending index of longest bitonic subarray
    end := 0
    // length of longest bitonic subarray
    maxlen := 1
    // array index
    i := 0
    // for each element excluding last
    for i + 1 < len(nums) {
        // length of current subarray
        curlen := 1
        // while subarray increasing
        for i + 1 < len(nums) && nums[i] < nums[i + 1] {
            // move to next element
            i++
            // increase subarray length
            curlen++
        }
        // while subarray decreasing
        for i + 1 < len(nums) && nums[i] > nums[i + 1] {
            // move to next element
            i++
            // increase subarray length
            curlen++
        }
        // if subarray length greater than current maximum
        if curlen > maxlen {
            // update maximum length
            maxlen = curlen
            // update ending index
            end = i
        }
    }
    // return subarray indices
    return end - maxlen + 1, end + 1
}

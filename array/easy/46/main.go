package main

import "fmt"

// program to find longest alternating subarray
func main() {
    // array of integers
    nums := []int{1, -2, 6, 4, -3, 2, -4, -3}
    // get indices of longest alternating subarray
    i, j := maxalter(nums)
    // print result
    fmt.Printf("%v\n", nums[i:j])
}

// find longest alternating subarray
func maxalter(nums []int) (int, int) {
    // length of current longest alternating subarray
    maxlen := 1
    // index of end of subarray
    end := 0
    // length of current subarray ending at current position
    curlen := 1
    // for each element past first
    for i := 1; i < len(nums); i++ {
        // if current sign is not previous sign
        if nums[i] * nums[i - 1] < 0 {
            // include current element in subarray
            curlen++
            // if current length is greater than max
            if curlen > maxlen {
                // update maximum
                maxlen = curlen
                // update ending index
                end = i
            }
        } else {
            // reset current length
            curlen = 1
        }
    }
    // return indices
    return end - maxlen + 1, end + 1
}

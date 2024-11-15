package main

import "fmt"

// program to find largest subarray with equal number of 0's and 1's
func main() {
    // array of integers
    nums := []int{0, 0, 1, 0, 1, 0, 0}
    // get subarray
    i, j := maxsub(nums)
    // print result
    fmt.Printf("%v\n", nums[i:j])
}

// find largest subarray with equal number of 0's and 1's
func maxsub(nums []int) (int, int) {
    // stores ending index of subarray with sum
    end := map[int]int{0: -1}
    // maximum length subarray
    maxlen := 0
    // ending index of maximum length subarray
    endidx := -1
    // current sum
    cursum := 0
    // for each element
    for i, v := range nums {
        // add element to sum (replace 0 with -1)
        if v == 0 {
            // add negative one
            cursum += -1
        } else {
            // or one
            cursum++
        }
        // if sum seen before
        if j, ok := end[cursum]; ok {
            // if length greater than maxlen
            if maxlen < i - j {
                // update maximum length
                maxlen = i - j
                // update ending index
                endidx = i
            }
        } else {
            // if sum seen for first time, insert mapping
            end[cursum] = i
        }
    }
    // return indices of subarray
    return (endidx - maxlen + 1), endidx + 1
}

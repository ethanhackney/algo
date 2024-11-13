package main

import "fmt"

// program to find maximum length subarray having a given sum
func main() {
    // array of integers
    nums := []int{5, 6, -5, 5, 3, 5, 3, -2, 0}
    // sum to look for
    sum := 8
    // find maximum length subarray
    i, j := maxlen(nums, sum)
    // print result
    fmt.Printf("%v\n", nums[i:j])
}

// find maximum length subarray having a given sum
func maxlen(nums []int, sum int) (int, int) {
    // stores ending index of subarray with given sum
    end := map[int]int{0: -1}
    // current sum
    cursum := 0
    // length of maximum length subarray
    maxlen := 0
    // ending index of maximum length subarray
    endidx := -1
    // for each element
    for i, v := range nums {
        // add current element to sum
        cursum += v
        // if seen first time, insert into map
        if _, ok := end[cursum]; !ok {
            // map current sum to current index
            end[cursum] = i
        }
        // if sum found
        if start, ok := end[cursum - sum]; ok {
            // if length greater than current maximum length
            if maxlen < i - start {
                // update maximum length
                maxlen = i - start
                // update ending index
                endidx = i
            }
        }
    }
    // return start and end indices of maximum length subarray
    return (endidx - maxlen + 1), endidx + 1
}

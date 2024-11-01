package main

import "fmt"

// program to find minimum index of repeating element in array
func main() {
    // array of integers
    nums := []int{5, 6, 3, 4, 3, 6, 4}
    // find minimum index of repeating element
    if i := minidx(nums); i != -1 {
        // print result
        fmt.Printf("%d\n", i)
    }
}

// find minimum index of repeating element
func minidx(nums []int) int {
    // index of minimum repeating element
    idx := len(nums)
    // keep track of elements we have seen
    seen := map[int]bool{}
    // for each element in reverse
    for i := len(nums) - 1; i >= 0; i-- {
        // if we have seen this element before
        if _, ok := seen[nums[i]]; ok {
            // update minimum index
            idx = i
        } else {
            // else, add nums[i] to seen elements
            seen[nums[i]] = true
        }
    }
    // if no repeating element
    if idx == len(nums) {
        // return no index
        return -1
    }
    // return index of repeating element
    return idx
}

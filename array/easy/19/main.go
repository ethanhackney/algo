package main

import (
    "fmt"
    "math"
)

// program to find index of element that is greater than all elements
// to the left and less than all elements to the right
func main() {
    // array of integers
    nums := []int{4, 2, 3, 5, 1, 6, 9, 7}
    // find index
    if i := index(nums); i != -1 {
        // print result
        fmt.Printf("%d\n", i)
    }
}

// find index of element that is greater than all elements to the left
// and less than all elements to the right
func index(nums []int) int {
    // stores max element in subarray nums[0..i-1]
    maxleft := make([]int, len(nums))
    // initialize maxleft[0] to minimum
    maxleft[0] = math.MinInt
    // for each element past the first
    for i := 1; i < len(nums); i++ {
        // save the maximum of previous element and this element
        maxleft[i] = max(maxleft[i - 1], nums[i - 1])
    }
    // store minimum value to the right
    curmin := nums[len(nums) - 1]
    // for each element from right to left
    for i := len(nums) - 2; i > 0; i-- {
        // if greater than left and less than right
        if maxleft[i] < nums[i] && nums[i] < curmin {
            // found index
            return i
        }
        // else update minimum value
        curmin = min(curmin, nums[i])
    }
    // no element satisfies constraint
    return -1
}

package main

import "fmt"

// program to find majority element in array
func main() {
    // array of integers
    nums := []int{1, 8, 7, 4, 1, 2, 2, 2, 2, 2, 2}
    // print majority element if there is one
    if i := majority(nums); i != -1 {
        fmt.Printf("%d\n", nums[i])
    }
}

// return majority element if there is one
func majority(nums []int) int {
    // index of majority element
    midx := -1
    // counter for current candidate
    count := 0
    // for each element
    for i, v := range nums {
        // if count is 0
        if count == 0 {
            // pick new candidate
            midx = i
            // reset counter
            count = 1
        } else {
            // increment if current value matches candidate
            if v == nums[midx] {
                count++
            } else {
                // otherwise, decrement count
                count--
            }
        }
    }
    // return majority element or -1 if none found
    return midx
}

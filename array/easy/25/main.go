package main

import "fmt"

// program to run insertion sort
func main() {
    // array of integers
    nums := []int{3, 8, 5, 4, 1, 9, -2}
    // run insertion sort
    isort(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// insertion sort
func isort(nums []int) {
    // for each element past the first
    for i := 1; i < len(nums); i++ {
        // save value because it may be overwritten
        v := nums[i]
        // index of current element
        j := i
        // while not at zero and v is less than nums[j - 1]
        for j > 0 && nums[j - 1] > v {
            // move up element
            nums[j] = nums[j - 1]
            // decrement index
            j--
        }
        // store v in its new position
        nums[j] = v
    }
}

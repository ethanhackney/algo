package main

import "fmt"

// program to move all zeros in an array to the end
func main() {
    // array of integers
    nums := []int{6, 0, 8, 2, 3, 0, 4, 0, 1}
    // move zeros
    move(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// move zeros to end of array
func move(nums []int) {
    // next place to put a nonzero value
    nonzero := 0
    // for each element
    for i := range nums {
        // if element is nonzero
        if nums[i] != 0 {
            // move into next nonzero place
            swap(&nums[i], &nums[nonzero])
            // move to next place
            nonzero++
        }
    }
}

// swap values at two pointers
func swap(ip, jp *int) {
    *ip, *jp = *jp, *ip
}

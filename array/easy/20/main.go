package main

import "fmt"

// program to left rotate an array
func main() {
    // array of integers
    nums := []int{1, 2, 3, 4, 5}
    // where to rotate
    pos := 2
    // rotate array
    rotate(nums, pos)
    // print result
    fmt.Printf("%v\n", nums)
}

// left rotate an array
func rotate(nums []int, k int) {
    // reverse first k elements
    rev(nums, 0, k - 1)
    // reverse n - k elements
    rev(nums, k, len(nums) - 1)
    // reverse all elements
    rev(nums, 0, len(nums) - 1)
}

// reverse portion of array
func rev(nums []int, i, j int) {
    // while pointers have not hit each other
    for i < j {
        // swap elements
        nums[i], nums[j] = nums[j], nums[i]
        // increment i
        i++
        // decrement j
        j--
    }
}

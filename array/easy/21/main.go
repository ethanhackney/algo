package main

import "fmt"

// program to right rotate an array
func main() {
    // array of integers
    nums := []int{1, 2, 3, 4, 5, 6, 7}
    // place to rotate on
    pos := 3
    // rotate array
    rotate(nums, pos)
    // print result
    fmt.Printf("%v\n", nums)
}

// right rotate an array
func rotate(nums []int, k int) {
    // reverse last k elements
    rev(nums, len(nums) - k, len(nums) - 1)
    // reverse first n - k elements
    rev(nums, 0, len(nums) - k - 1)
    // reverse whole array
    rev(nums, 0, len(nums) - 1)
}

// reverse portion of array
func rev(nums []int, i, j int) {
    // while i is less than j
    for i < j {
        // swap elements
        nums[i], nums[j] = nums[j], nums[i]
        // increment i
        i++
        // decrement j
        j--
    }
}

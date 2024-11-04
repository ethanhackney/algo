package main

import "fmt"

// program to segregate positive and negative integers in linear time
func main() {
    // array of integers
    nums := []int{9, -3, 5, -2, -8, -6, 1, 3}
    // segregate
    segregate(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// segregate positive and negative integers
func segregate(nums []int) {
    // index of next place to put a negative value
    neg := 0
    // for each element
    for i := range nums {
        // if negative
        if nums[i] < 0 {
            // add negative integer to negative subarray
            swap(nums, i, neg)
            // increment negative subarray pointer
            neg++
        }
    }
}

// swap elements of array
func swap(nums []int, i, j int) {
    nums[i], nums[j] = nums[j], nums[i]
}

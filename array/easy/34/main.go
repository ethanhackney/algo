package main

import "fmt"

// program to reverse an array
func main() {
    // array of integers
    nums := []int{1, 2, 3, 4, 5}
    // reverse array
    rev(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// reverse an array
func rev(nums []int) {
    // index that moves forward 
    i := 0
    // index that moves backward
    j := len(nums) - 1
    // while i is less than j
    for i < j {
        // swap elements
        swap(&nums[i], &nums[j])
        // move i up one
        i++
        // move j down one
        j--
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    // swap values
    *ip, *jp = *jp, *ip
}

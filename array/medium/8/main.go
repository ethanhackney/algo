package main

import "fmt"

// program to rearrange an array with alternating high and low elements
func main() {
    // array of integers
    nums := []int{9, 6, 8, 3, 7}
    // rearrange array
    rearrange(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// rearrange array
func rearrange(nums []int) {
    // for every other element
    for i := 1; i < len(nums); i += 2 {
        // if previous element is greater
        if nums[i - 1] > nums[i] {
            // swap current with previous
            swap(&nums[i - 1], &nums[i])
        }
        // if next element is greater
        if i + 1 < len(nums) && nums[i + 1] > nums[i] {
            // swap current with next
            swap(&nums[i + 1], &nums[i])
        }
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    // swap values
    *ip, *jp = *jp, *ip
}

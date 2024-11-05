package main

import "fmt"

// program to run bubble sort on array
func main() {
    // array of integers
    nums := []int{3, 5, 8, 4, 1, 9, -2}
    // run bubble sort
    bsort(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// run bubble sort
func bsort(nums []int) {
    // for each element except last
    for i := 0; i < len(nums) - 1; i++ {
        // last i elements are already sorted, so run 0 to n - i - 1 times
        for j := 0; j < len(nums) - i - 1; j++ {
            // if current element greater than next
            if nums[j] > nums[j + 1] {
                // swap elements
                swap(&nums[j], &nums[j + 1])
            }
        }
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    *ip, *jp = *jp, *ip
}

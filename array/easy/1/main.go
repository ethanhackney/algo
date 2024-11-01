package main

import "fmt"

// program to sort a binary array in linear time
func main() {
    // array of integers
    nums := []int{1, 0, 0, 0, 1, 0, 1, 1}
    // sort array
    sort(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// sort binary array
func sort(nums []int) {
    // pivot value
    pivot := 1
    // index of next unsorted value
    unsorted := 0
    // for each element
    for i := range nums {
        // if less than pivot
        if nums[i] < pivot {
            // swap with next unsorted value
            swap(&nums[i], &nums[unsorted])
            // advance to next unsorted value
            unsorted++
        }
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    *ip, *jp = *jp, *ip
}

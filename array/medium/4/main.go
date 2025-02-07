package main

import "fmt"

// program to sort an array of 0's, 1's, and 2's
func main() {
    // array of 0's, 1's, and 2's
    nums := []int{0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0}
    // sort array
    sort(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// sort an array of 0's, 1's, and 2's
func sort(nums []int) {
    // low pointer of window
    low := 0
    // middle pointer of window
    mid := 0
    // pivot element
    pivot := 1
    // high pointer of window
    high := len(nums) - 1
    // while window non-empty
    for mid <= high {
        // if element is zero
        if nums[mid] < pivot {
            // swap with low element
            swap(&nums[low], &nums[mid])
            // advance low pointer
            low++
            // advance mid pointer
            mid++
        } else if nums[mid] > pivot {
            // if element is 2
            // swap with mid element
            swap(&nums[mid], &nums[high])
            // move high pointer down
            high--
        } else {
            // if element is 1
            // advance mid pointer
            mid++
        }
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    // swap values
    *ip, *jp = *jp, *ip
}

package main

import "fmt"

// program to sort an array with two misplaced elements
func main() {
    // array of integers
    nums := []int{3, 5, 6, 9, 8, 7}
    // sort array
    sort(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// sort an array of integers with two misplaced elements
func sort(nums []int) {
    // index of first misplaced element
    first := -1
    // index of second misplaced element
    second := -1
    // previous element
    prev := nums[0]
    // for every element past the first
    for i := 1; i < len(nums); i++ {
        // if previous element is greater than current
        if prev > nums[i] {
            // if first conflict
            if first == -1 {
                // update first index
                first = i - 1
            }
            // second occurence
            second = i
        }
        // update previous element
        prev = nums[i]
    }
    // swap misplaced elements
    swap(&nums[first], &nums[second])
}

// swap values at two pointers
func swap(ip, jp *int) {
    *ip, *jp = *jp, *ip
}

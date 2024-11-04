package main

import "fmt"

// program to run selection sort on array
func main() {
    // array of integers
    nums := []int{3, 5, 8, 4, 1, 9, -2}
    // run selection sort
    ssort(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// selection sort
func ssort(nums []int) {
    // for each element
    for i := 0; i < len(nums) - 1; i++ {
        // start with current element as minimum
        min := i
        // for rest of elements
        for j := i + 1; j < len(nums); j++ {
            // if current element is less than minimum
            if nums[j] < nums[min] {
                // update minimum
                min = j
            }
        }
        // swap minimum element into place
        swap(&nums[min], &nums[i])
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    *ip, *jp = *jp, *ip
}

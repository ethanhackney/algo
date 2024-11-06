package main

import "fmt"

// program to sort binary array in linear time and constant space
func main() {
    // binary array
    nums := []int{1, 0, 0, 0, 1, 0, 1, 1}
    // sort array
    sort(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// sort binary array
func sort(nums []int) {
    // next place to write a zero
    zero := 0
    // for each element
    for i := range nums {
        // if nums[i] is zero
        if nums[i] == 0 {
            // move zero into place
            swap(&nums[i], &nums[zero])
            // move to next place to write a zero
            zero++
        }
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    *ip, *jp = *jp, *ip
}

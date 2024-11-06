package main

import "fmt"

// program to rearrange array so positive and negative values
// occupy alternate indices
func main() {
    // array of integers
    nums := []int{9, -3, 5, -2, -8, -6, 1, 3}
    // rearrange array
    rearrange(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// rearrange array so positive and negative values occupy alternate indices
func rearrange(nums []int) {
    // partition array into positive and negative subarrays
    p := part(nums)
    // while more negative or positive elements
    for i := 0; i < len(nums); i += 2 {
        // if no more positive elements
        if i == p {
            // break out of loop
            break
        }
        // swap elements
        swap(&nums[i], &nums[p])
        // move to next positive element
        p++
    }
}

// move positive values together and negative values together
func part(nums []int) int {
    // next place to write a negative value to
    neg := 0
    // for each element
    for i := range nums {
        if nums[i] < 0 {
            // move negative value into place
            swap(&nums[i], &nums[neg])
            // move to next place to write negative value
            neg++
        }
    }
    // return one past last negative value
    return neg
}

// swap values of two pointers
func swap(ip, jp *int) {
    *ip, *jp = *jp, *ip
}

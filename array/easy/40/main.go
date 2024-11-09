package main

import "fmt"

// program to find floor and ceil of a number in a sorted array
func main() {
    // array of integers
    nums := []int{1, 4, 6, 8, 9}
    // number to find floor and ceil for
    v := 5
    // print floor
    fmt.Printf("%d\n", nums[floor(nums, v)])
    // print ceil
    fmt.Printf("%d\n", nums[ceil(nums, v)])
}

// find floor of a number in a sorted array
func floor(nums []int, v int) int {
    // low pointer of window
    low := 0
    // high pointer of window
    high := len(nums) - 1
    // index of floor
    fidx := -1
    // while window non-empty
    for low <= high {
        // get mid point
        mid := low + ((high - low) / 2)
        // if v found, return mid
        if v == nums[mid] {
            return mid
        } else if v < nums[mid] {
            // if less than mid, move high pointer down
            high = mid - 1
        } else {
            // if greater than mid, update floor
            fidx = mid
            // and move low pointer up
            low = mid + 1
        }
    }
    // return index of floor
    return fidx
}

// find ceil of a number in a sorted array
func ceil(nums []int, v int) int {
    // low pointer of window
    low := 0
    // high pointer of window
    high := len(nums) - 1
    // index of ceil
    cidx := -1
    // while window non-empty
    for low <= high {
        // get mid point
        mid := low + ((high - low) / 2)
        // if v found, return mid
        if v == nums[mid] {
            return mid
        } else if v < nums[mid] {
            // if less than mid, update ceil
            cidx = mid
            // and move high pointer down
            high = mid - 1
        } else {
            // if greater than mid, move low pointer up
            low = mid + 1
        }
    }
    // return index of ceil
    return cidx
}

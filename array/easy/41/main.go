package main

import "fmt"

// program to find the number of 1's in a sorted binary array
func main() {
    // sorted binary array
    nums := []int{0, 0, 0, 0, 1, 1, 1}
    // print result
    fmt.Printf("%d\n", count(nums, 0, len(nums) - 1))
}

// find number of 1's in a sorted binary array
func count(nums []int, low, high int) int {
    // if last element is 0, no 1's are present
    if nums[high] == 0 {
        // count of 1's is zero
        return 0
    }
    // if first element is 1, all elements are 1
    if nums[low] == 1 {
        // return subarray count
        return high - low + 1
    }
    // get mid point
    mid := low + ((high - low) / 2)
    // divide array and recur for each subarray
    return count(nums, low, mid) + count(nums, mid + 1, high)
}

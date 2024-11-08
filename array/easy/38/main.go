package main

import "fmt"

// program to count number of rotations in circularly sorted array
func main() {
    // array of integers
    nums := []int{8, 9, 10, 2, 5, 6}
    // print result
    fmt.Printf("%d\n", count(nums))
}

// count number of rotations in circularly sorted array
func count(nums []int) int {
    // low pointer of window
    low := 0
    // high pointer to window
    high := len(nums) - 1
    // while window is non-empty
    for low <= high {
        // if window already sorted
        if nums[low] <= nums[high] {
            // found minimum element
            return low
        }
        // get mid point
        mid := low + ((high - low) / 2)
        // next element
        next := (mid + 1) % len(nums)
        // previous element
        prev := (mid - 1 + len(nums)) % len(nums)
        // if middle element is less than next and prev
        if nums[mid] <= nums[next] && nums[mid] <= nums[prev] {
            // found minimum element
            return mid
        } else if nums[mid] <= nums[high] {
            // if nums[mid] <= nums[high], move high pointer down
            high = mid - 1
        } else {
            // if nums[mid] >= nums[low], move low pointer up
            low = mid + 1
        }
    }
    // array not circularly sorted
    return -1
}

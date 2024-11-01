package main

import "fmt"

// program to partition an array into two subarrays with same sum
func main() {
    // array of integers
    nums := []int{6, -4, -3, 2, 3}
    // if we can partition
    if i := part(nums); i != -1 {
        // print first subarray
        fmt.Printf("%v\n", nums[0:i])
        // print second subarray
        fmt.Printf("%v\n", nums[i:])
    }
}

// partition array into two subarrays with same sum
func part(nums []int) int {
    // calculate sum of elements
    allsum := sum(nums)
    // maintain current sum
    cursum := 0
    // for each element
    for i := range nums {
        // if sum of nums[0...i-1] is equal to sum of nums[i,n-1] 
        if cursum == allsum - cursum {
            // return index of place to partition array
            return i
        }
        // add current element to sum
        cursum += nums[i]
    }
    // index to partition on not found
    return -1
}

// sum up array elements
func sum(nums []int) int {
    // holds sum
    sum := 0
    // for each element
    for _, v := range nums {
        // add v to sum
        sum += v
    }
    // return sum
    return sum
}

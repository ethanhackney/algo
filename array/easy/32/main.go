package main

import "fmt"

// program to find missing number in an array
func main() {
    // array of integers
    nums := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
    // print missing number
    fmt.Printf("%v\n", missing(nums))
}

// find missing number in an array
func missing(nums []int) int {
    // xor of all elements
    xor := 0
    // for each element
    for _, v := range nums {
        // add v to xor
        xor ^= v
    }
    // compute xor from 1 to n + 1
    for n := 1; n <= len(nums) + 1; n++ {
        // add n to xor
        xor ^= n
    }
    // return missing number
    return xor
}

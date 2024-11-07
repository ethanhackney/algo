package main

import "fmt"

// program to find odd occurring element in an array
func main() {
    // array of integers
    nums := []int{4, 3, 6, 2, 6, 4, 2, 3, 4, 3, 3}
    // print result
    fmt.Printf("%d\n", odd(nums))
}

// find odd occurring element in an array
func odd(nums []int) int {
    // xor of all elements
    xor := 0
    // for each element
    for _, v := range nums {
        // add v to xor
        xor ^= v
    }
    // return odd number
    return xor
}

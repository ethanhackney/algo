package main

import "fmt"

// program to set both elements of binary array to 0 in a single line
func main() {
    // binary array
    nums := []int{1, 0}
    // do conversion
    convert(nums)
    // print result
    fmt.Printf("%v\n", nums)
}

// set both elements of binary array to 0 in a single line
func convert(nums []int) {
    // set both elements to zero
    nums[nums[1]] = nums[nums[0]]
}

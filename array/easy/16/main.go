package main

import "fmt"

// program to add two arrays into a new array
func main() {
    // two arrays of integers
    nums := [][]int{
        { 23,  5, 2, 7, 87 },
        {  4, 67, 2, 8     },
    }
    // add arrays
    digits := add(nums[0], nums[1])
    // print result
    fmt.Printf("%v\n", digits)
}

// add two arrays into a new array
func add(a, b []int) []int {
    // length of a
    m := len(a)
    // length of b
    n := len(b)
    // holds digits of result
    digits := []int{}
    // index into a and b
    i := 0
    // while a and b are not exhausted
    for i < m && i < n {
        // get sum of array elements
        sum := a[i] + b[i]
        // split sum into digits
        digits = split(digits, sum)
        // move to next pair
        i++
    }
    // add rest of a elements if any
    for i < m {
        // split a[i] into digits
        digits = split(digits, a[i])
        // move to next element
        i++
    }
    // add rest of b elements if any
    for i < n {
        // split a[i] into digits
        digits = split(digits, b[i])
        // move to next element
        i++
    }
    // return digits
    return digits
}

// add digits of number into array
func split(digits []int, n int) []int {
    // if more digits
    if n > 0 {
        // recurse for rest of digits
        digits = split(digits, n / 10)
        // add current digit to digits
        digits = append(digits, n % 10)
    }
    // return result
    return digits
}

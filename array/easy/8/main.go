package main

import (
    "fmt"
    "math"
)

// program to find maximum product subset
func main() {
    // array of integers
    nums := []int{-6, 4, -5, 8, -10, 0, 8}
    // print maximum product
    pidx := maxprod(nums)
    // calculate maximum product
    prod := 1
    // for each element in pidx
    for _, idx := range pidx {
        // multiply prod by nums[idx]
        prod *= nums[idx]
    }
    // print maximum product
    fmt.Printf("%d\n", prod)
}

// find maximum product subset
func maxprod(nums []int) []int {
    // indices of elements in maximum product subset
    prod := []int{}
    // stores negative element with minimum absolute value
    absmin := math.MaxInt
    // index of element with minimum absolute value
    absidx := -1
    // count of negative elements
    neg := 0
    // count of positive elements
    pos := 0
    // flag says whether a zero was found or not
    zero := false
    // indices of zeros
    zeroidx := []int{}
    // for each element
    for i, v := range nums {
        // if negative
        if v < 0 {
            // increment negative count
            neg++
            // update absmin and absidx if absolute value is less than it
            if abs(v) < absmin {
                // update minimum absolute value
                absmin = abs(v)
                // update index of element with minimum absolute value
                absidx = i
            }
        } else if v > 0 {
            // increment positive count
            pos++
        }
        // if zero is found
        if v == 0 {
            // set zero flag
            zero = true
            // add index of zero to zero list
            zeroidx = append(zeroidx, i)
        } else {
            // add element to product
            prod = append(prod, i)
        }
    }
    // if odd number of negatives
    if (neg & 1) != 0 {
        // exclude negative element with minimum absolute value
        prod = append(prod[:absidx], prod[absidx+1:]...)
    }
    // set contains one negative element and one or more zeros
    if neg == 1 && pos == 0 && zero {
        // set equal to list of zeros
        prod = zeroidx
    }
    // set contains all 0s
    if neg == 0 && pos == 0 && zero {
        // set equal to list of zeros
        prod = zeroidx
    }
    // return maximum produt
    return prod
}

// return absolute value of integer
func abs(n int) int {
    // if negative
    if n < 0 {
        return -n
    }
    // if positive
    return n
}

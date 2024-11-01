package main

import (
    "fmt"
    "math"
)

// program to find minimum difference between indices of two elements in array
func main() {
    // array of integers
    nums := []int{1, 3, 5, 4, 8, 2, 4, 3, 6, 5}
    // x value
    x := 2
    // y value
    y := 5
    // print indices
    i, j := mindiff(nums, x, y)
    // print minimum difference
    fmt.Printf("%d\n", abs(i - j))
}

// find minimum difference between indices of two elements
func mindiff(nums []int, x, y int) (int, int) {
    // minimum x index
    xminidx := -1
    // minimum y index
    yminidx := -1
    // index of x value
    xidx := -1
    // index of y value
    yidx := -1
    // minimum difference between indices of x and y
    mindiff := math.MaxInt
    // for each element
    for i, v := range nums {
        // if v is x
        if v == x {
            // update index of x
            xidx = i
            // if we have seen y
            if yidx != -1 {
                // calculate difference
                diff := abs(xidx - yidx)
                // if difference is less than minimum
                if diff < mindiff {
                    // update minimum difference
                    mindiff = diff
                    // update minimum x index
                    xminidx = xidx
                    // update minimum y index
                    yminidx = yidx
                }
            }
        }
        // if v is y
        if v == y {
            // update index of y
            yidx = i
            // if we have seen x
            if xidx != -1 {
                // calculate difference
                diff := abs(xidx - yidx)
                // if difference is less than minimum
                if diff < mindiff {
                    // update minimum difference
                    mindiff = diff
                    // update minimum x index
                    xminidx = xidx
                    // update minimum y index
                    yminidx = yidx
                }
            }
        }
    }
    // return indices
    return xminidx, yminidx
}

// return absolute value of integer
func abs(n int) int {
    // if negative
    if n < 0 {
        // negate it
        n = -n
    }
    // return absolute value
    return n
}

package main

import (
    "fmt"
    "math"
)

// program to determine whether a BST is skewed from its preorder traversal
func main() {
    // preorder traversal
    pre := []int{15, 30, 25, 18, 20}
    // print result
    fmt.Printf("%v\n", skewed(pre))
}

// determine whether a BST is skewed from its preorder traversal
func skewed(pre []int) bool {
    // minimum value of range
    min := math.MinInt
    // maximum value of range
    max := math.MaxInt
    // for every element past first
    for i := 1; i < len(pre); i++ {
        // if current element is in range
        if pre[i] >= min && pre[i] <= max {
            // update valid range for next element
            if pre[i] > pre[i - 1] {
                // update minimum
                min = pre[i - 1] + 1
            } else {
                // update maximum
                max = pre[i - 1] - 1
            }
        } else {
            // BST is not skewed
            return false
        }
    }
    // BST is skewed
    return true
}

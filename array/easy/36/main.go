package main

import "fmt"

// program to print complete BST in increasing order
func main() {
    // array of integers
    nums := []int{15, 10, 20, 8, 12, 18, 25}
    // print result
    for _, v := range increasing(nums) {
        // print node
        fmt.Printf("%d ", nums[v])
    }
    // print newline
    fmt.Printf("\n")
}

// return complete BST in increasing order
func increasing(nums []int) []int {
    // holds elements of BST in increasing order
    order := []int{}
    // stack of indices
    stk := []int{}
    // set curr equal to root node
    curr := 0
    // push root node onto stack
    stk = append(stk, curr)
    // while more elements in stack
    for len(stk) > 0 {
        // descend to the left
        for curr == stk[len(stk) - 1] && curr * 2 + 1 < len(nums) {
            // move to left child
            curr = curr * 2 + 1
            // add index to stack
            stk = append(stk, curr)
        }
        // pop top element
        curr = stk[len(stk) - 1]
        stk = stk[:len(stk) - 1]
        // add top to order
        order = append(order, curr)
        // descend to the right
        if curr * 2 + 2 < len(nums) {
            // move to the right
            curr = curr * 2 + 2
            // push right child
            stk = append(stk, curr)
        }
    }
    // return elements in increasing order
    return order
}

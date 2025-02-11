package main

import "fmt"

// program to replace every array element with the product of every other
// element without using a division operator
func main() {
    // array of integers
    nums := []int{5, 3, 4, 2, 6, 8}
    // do replacement
    replace(nums, 1, 0)
    // print result
    fmt.Printf("%v\n", nums)
}

// replace every array element with the product of every other element
func replace(nums []int, left, i int) int {
    // base case: no elements left
    if i == len(nums) {
        // base case is product of 1
        return 1
    }
    // save value of current element
    cur := nums[i]
    // get product of right subarray
    right := replace(nums, left * nums[i], i + 1)
    // replace current element with product of right subarray
    nums[i] = left * right
    // return product of current element and right subarray product
    return cur * right
}

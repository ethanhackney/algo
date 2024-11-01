package main

import "fmt"

// program to find integers that sum to a value in an array
func main() {
    // array of integers
    nums := []int{8, 7, 2, 5, 3, 1}
    // sum we are looking for
    sum := 10
    // if any pairs were found
    if p := pairs(nums, sum); len(p) != 0 {
        // print pairs
        for _, v := range p {
            // print pair
            fmt.Printf("%d, %d\n", nums[v[0]], nums[v[1]])
        }
    }
}

// find pairs that sum to sum in an array of distinct integers
func pairs(nums []int, sum int) [][]int {
    // maps values to indices
    idx := map[int]int{}
    // array of pairs
    pairs := [][]int{}
    // for each element
    for i, v := range nums {
        // if we have seen (sum - v), then we found a pair
        if j, ok := idx[sum - v]; ok {
            // add pair to pairs
            pairs = append(pairs, []int{j, i})
        }
        // save index of value
        idx[v] = i
    }
    // return pairs found
    return pairs
}

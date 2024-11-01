package main

import "fmt"

// program to find all pairs with difference k
func main() {
    // array of integers
    nums := []int{1, 5, 2, 2, 2, 5, 5, 4}
    // difference to search for
    k := 3
    // print pairs with difference k
    fmt.Printf("%v\n", pairs(nums, k))
}

// find indices of all pairs with difference k
func pairs(nums []int, k int) [][]int {
    // mapping from value to indices where value occurs
    idx := map[int][]int{}
    // add indices to map
    for i, v := range nums {
        // add index to idx[v]
        idx[v] = append(idx[v], i)
    }
    // keep track of which pairs we have already added to pairs
    seen := make([][]bool, len(nums) + 1)
    // add rows to seen matrix
    for i := range seen {
        // add row
        seen[i] = make([]bool, len(nums) + 1)
    }
    // list of pairs
    pairs := [][]int{}
    // for each element
    for i, v := range nums {
        // if v - k is found
        if _, ok := idx[v - k]; ok {
            // add pairs
            for _, j := range idx[v - k] {
                // if we have not seen this pair before
                if !seen[i][j] && !seen[j][i] {
                    // add pair
                    pairs = append(pairs, []int{j, i})
                    // mark pair as seen
                    seen[i][j] = true
                    seen[j][i] = true
                }
            }
        }
        // if v + k is found
        if _, ok := idx[v + k]; ok {
            // add pairs
            for _, j := range idx[v + k] {
                // if we have no seen this pair before
                if !seen[i][j] && !seen[j][i] {
                    // add pair
                    pairs = append(pairs, []int{j, i})
                    // mark pair as seen
                    seen[i][j] = true
                    seen[j][i] = true
                }
            }
        }
    }
    // return pairs found
    return pairs
}

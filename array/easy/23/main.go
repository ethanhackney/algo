package main

import "fmt"

// program to count triplets which form an inversion in an array
func main() {
    // array of integers
    nums := []int{9, 4, 3, 5, 1}
    // print inversion count
    fmt.Printf("%v\n", inversion(nums))
}

// count triplets which form an inversion in an array
func inversion(nums []int) [][]int {
    // array of triplet inversions
    inv := [][]int{}
    // for each element from 0 to until len(nums) - 2
    for i := 0; i < len(nums) - 2; i++ {
        // for each element from i + 1 to len(nums) - 1
        for j := i + 1; j < len(nums) - 1; j++ {
            // for each element from j + 1 to len(nums)
            for k := j + 1; k < len(nums); k++ {
                // if i, j, k is an inversion
                if nums[i] > nums[j] && nums[j] > nums[k] {
                    // add inversion
                    inv = append(inv, []int{i, j, k})
                }
            }
        }
    }
    // return inversions
    return inv
}

package main

import "fmt"

// program to print all subarrays with 0 sum
func main() {
    // array of integers
    nums := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
    // print result
    fmt.Printf("%v\n", zerosum(nums))
}

// print all subarrays with 0 sum
func zerosum(nums []int) [][]int {
    // stores subarrays
    subs := [][]int{}
    // store ending index of all subarrays having same sum
    end := map[int][]int{}
    // initial element for zero-sum subarrays starting at 0
    end[0] = append(end[0], -1)
    // running sum
    sum := 0
    // for each element
    for i, v := range nums {
        // add element so sum
        sum += v
        // if seen sum before
        if idx, ok := end[sum]; ok {
            // add subarrays
            for _, v := range idx {
                // add subarray
                subs = append(subs, []int{v + 1, i})
            }
        }
        // add sums ending index
        end[sum] = append(end[sum], i)
    }
    // return subarrays
    return subs
}

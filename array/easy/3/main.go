package main

import "fmt"

// program to find indices of equilibrium values in an array
func main() {
    // array of integers
    nums := []int{0, -3, 5, -4, -2, 3, 1, 0}
    // print equilibriums
    fmt.Printf("%v\n", equilibrium(nums))
}

func equilibrium(nums []int) []int {
    // indices of equilibriums
    eq := []int{}
    // get sum of all elements
    sum := sum(nums)
    // sum of elements in nums[i+1...n)
    cursum := 0
    // for each element from last to first
    for i := len(nums) - 1; i >= 0; i-- {
        // is equilibrium if sum of nums[0...i-1] is equal to nums[i+1...n)
        if cursum == sum - (nums[i] + cursum) {
            // add to array of equilibriums
            eq = append(eq, i)
        }
        // add current element to sum
        cursum += nums[i]
    }
    // return found equilibriums
    return eq
}

// sum up elements of array
func sum(nums []int) int {
    // where to store sum
    sum := 0
    // add all elements to sum
    for _, v := range nums {
        // add element to sum
        sum += v
    }
    // return sum
    return sum
}

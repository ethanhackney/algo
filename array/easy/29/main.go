package main

import "fmt"

// program to run counting sort on array
func main() {
    // array of integers
    nums := []int{4, 2, 10, 10, 1, 4, 2, 1, 10}
    // maximum value in nums
    max := 10
    // run counting sort
    csort(nums, max)
    // print result
    fmt.Printf("%v\n", nums)
}

// run counting sort
func csort(nums []int, max int) {
    // mapping of value to count
    count := make([]int, max + 1)
    // store counts of each element
    for _, v := range nums {
        // increment v count
        count[v]++
    }
    // calculate start index of each value
    i := 0
    // for each integer from 0 to max inclusive
    for j := 0; j <= max; j++ {
        // save count in temporary
        tmp := count[j]
        // set starting index
        count[j] = i
        // increment tmp positions
        i += tmp
    }
    // auxiliary array
    aux := make([]int, len(nums))
    // copy into auxiliary array
    for _, v := range nums {
        // add element to auxiliary array
        aux[count[v]] = v
        // increment count
        count[v]++
    }
    // copy back into nums
    for i := range nums {
        // copy element into nums
        nums[i] = aux[i]
    }
}

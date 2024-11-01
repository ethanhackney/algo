package main

import (
    "fmt"
    "math/rand"
    "time"
)

// program to find index of maximum occurring element with equal probability
func main() {
    // array of integers
    nums := []int{4, 3, 6, 8, 4, 6, 2, 4, 5, 9, 7, 4}
    // print maximum occurring element 5 times
    for i := 0; i < 5; i++ {
        // print result
        fmt.Printf("%d\n", index(nums))
    }
}

func index(nums []int) int {
    // mapping of value to count
    count := map[int]int{}
    // for each element
    for _, v := range nums {
        // increment count
        count[v]++
    }
    // find first maximum occurring element
    maxcount := nums[0]
    // for each element
    for _, v := range nums {
        // if this element occurs more often than maxcount, update maxcount
        maxcount = max(maxcount, count[v])
    }
    // seed the random number generator
    rand.Seed(time.Now().UnixNano())
    // generate random number
    k := (rand.Int() % count[maxcount]) + 1
    // until we have found k'th element or end of array
    i := 0
    for k > 0 && i < len(nums) {
        // if equal to max occurring element
        if nums[i] == maxcount {
            // decrement k
            k--
        }
        // increment index
        i++
    }
    // return k'th index
    return i - 1
}

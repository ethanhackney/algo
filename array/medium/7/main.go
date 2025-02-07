package main

import (
    "fmt"
    "math/rand"
    "time"
)

// program to shuffle array using Fisher-Yates shuffle algorithm
func main() {
    // array of integers
    nums := []int{1, 2, 3, 4, 5, 6}
    // shuffle array
    shuffle(nums)
    // print array
    fmt.Printf("%v\n", nums)
}

// shuffle array using Fisher-Yates shuffle algorithm
func shuffle(nums []int) {
    // seed random number generator
    rand.Seed(time.Now().UnixNano())
    // for each middle element
    for i := len(nums) - 1; i >= 1; i-- {
        // generate random number in between 0 and i
        j := rand.Int() % (i + 1)
        // swap current element with element at j
        swap(&nums[i], &nums[j])
    }
}

// swap values of two pointers
func swap(ip, jp *int) {
    // swap values
    *ip, *jp = *jp, *ip
}

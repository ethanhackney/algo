package main

import "fmt"

// program to convert max heap to min heap in linear time
func main() {
    // max heap of integers
    maxheap := []int{9, 4, 7, 1, -2, 6, 5}
    // do conversion
    maxtomin(maxheap)
    // print result
    fmt.Printf("%v\n", maxheap)
}

// convert max heap to min heap
func maxtomin(nums []int) {
    // start at last internal node
    i := (len(nums) - 2) / 2
    // up to and including root node
    for i >= 0 {
        // heapify down on internal node
        down(nums, i)
        // move up tree
        i--
    }
}

// heapify down
func down(nums []int, i int) {
    // forever
    for {
        // index of minimum element in subtree
        min := i
        // left child
        l := left(i)
        // right child
        r := right(i)
        // if left child is less than root
        if l < len(nums) && nums[l] < nums[i] {
            // left child becomes minimum
            min = l
        }
        // if right child is less than min
        if r < len(nums) && nums[r] < nums[min] {
            // right child becomes minimum
            min = r
        }
        // if root is minimum
        if min == i {
            // we are done
            break
        }
        // swap root with minimum
        swap(&nums[i], &nums[min])
        // move down tree
        i = min
    }
}

// get left child
func left(i int) int {
    // left child is at index (i * 2) + 1
    return (i * 2) + 1
}

// get right child
func right(i int) int {
    // right child is at index (i * 2) + 2
    return (i * 2) + 2
}

// swap values of two pointers
func swap(ip, jp *int) {
    // swap values
    *ip, *jp = *jp, *ip
}

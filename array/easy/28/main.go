package main

import "fmt"

// program to run merge sort on array
func main() {
    // array of integers
    nums := []int{5, 4, 3, 2, 1}
    // auxiliary array
    aux := make([]int, len(nums))
    // copy array elements into auxiliary array
    copy(aux, nums)
    // sort array
    msort(nums, aux, 0, len(nums) - 1)
    // print result
    fmt.Printf("%v\n", nums)
}

// run merge sort
func msort(nums, aux []int, low, high int) {
    // if subarray length greater than zero
    if low < high {
        // find midpoint
        mid := low + ((high - low) / 2)
        // sort left subarray
        msort(nums, aux, low, mid)
        // sort right subarray
        msort(nums, aux, mid + 1, high)
        // merge subarrays
        merge(nums, aux, low, mid, high)
    }
}

// merge two subarrays
func merge(nums, aux []int, low, mid, high int) {
    // index of next place to write to in auxiliary array
    k := low
    // index into left subarray
    i := low
    // index into right subarray
    j := mid + 1
    // while more elements in each subarray
    for i <= mid && j <= high {
        // if left subarray element <= to right subarray element
        if nums[i] <= nums[j] {
            // add left subarray element to auxiliary array
            aux[k] = nums[i]
            // move to next element in left subarray
            i++
        } else {
            // add right subarray element to auxiliary array
            aux[k] = nums[j]
            // move to next element in right subarray
            j++
        }
        // move to next place to write to in auxiliary array
        k++
    }
    // copy remaining elements
    for i <= mid {
        // copy element to auxiliary array
        aux[k] = nums[i]
        // move to next place to write to in auxiliary array
        k++
        // move to next element in left subarray
        i++
    }
    // copy contents of auxiliary array back into nums
    for i := low; i <= high; i++ {
        // copy element to nums
        nums[i] = aux[i]
    }
}

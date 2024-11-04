package main

import (
    "fmt"
    "sort"
)

// represents an activity with a start and finish time
type Activity struct {
    Start  int // when activity starts
    Finish int // when activity finishes
}

// array of activities
type ActivityList []Activity

// get length of activity list
func (a ActivityList) Len() int {
    return len(a)
}

// swap elements in activity list
func (a ActivityList) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

// compare activities
func (a ActivityList) Less(i, j int) bool {
    return a[i].Finish < a[j].Finish
}

// program to find maximum number of activities that can be performed
func main() {
    // array of activities
    act := []Activity{
        {1,   4},
        {3,   5},
        {0,   6},
        {5,   7},
        {3,   8},
        {5,   9},
        {6,  10},
        {8,  11},
        {8,  12},
        {2,  13},
        {12, 14},
    }
    // print selected activities
    fmt.Printf("%v\n", selectActivity(act))
}

// find maximum number of activities that can be performed
func selectActivity(act ActivityList) []int {
    // keep track of last selected activity
    last := 0
    // stores indices of selected activities
    idx := []int{
        // store first index
        0,
    }
    // sort activities by finishing time
    sort.Sort(act)
    // for each element after the first
    for i := 1; i < act.Len(); i++ {
        // if start time is greater than or equal to last finish time
        if act[i].Start >= act[last].Finish {
            // add index to selected activities
            idx = append(idx, i)
            // update last selected activity
            last = i
        }
    }
    // return selected activities
    return idx
}

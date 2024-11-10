package main

import "fmt"

// program to find maximum profit that can be earned by selling stocks
func main() {
    // stock prices
    prices := [][]int{
        {5, 3, 4, 6,  3},
        {8, 4, 3, 5, 10},
    }
    // print result
    fmt.Printf("%d\n", maxprofit(prices))
}

// find maximum profit that can be earned by selling stocks
func maxprofit(prices [][]int) int {
    // previous price of previous price
    prevprev := 0
    // previous price
    prev := max(prices[0][0], prices[1][0])
    // for each element past first
    for i := 1; i < len(prices[0]); i++ {
        // current price
        curr := max(prices[0][i] + prev, prices[1][i] + prevprev)
        // update previous previous
        prevprev = prev
        // update prev
        prev = curr
    }
    // prev holds maximum profit
    return prev
}

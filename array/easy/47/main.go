package main

import "fmt"

// program to find itinerary from given list of departure and arrival airports
func main() {
    tickets := [][]string{
        {"LAX", "DXB"},
        {"DFW", "JFK"},
        {"LHR", "DFW"},
        {"JFK", "LAX"},
    }
    // print result
    for _, v := range itinerary(tickets) {
        // print source and destination airport
        fmt.Printf("%s -> %s\n", v[0], v[1])
    }
}

// find itinerary from given list of departure and arrival airports
func itinerary(tickets [][]string) [][]string {
    // mapping from departure to arrival airports
    tickmap := map[string]string{}
    // for each ticket
    for _, v := range tickets {
        // add mapping from departure to arrival airport
        tickmap[v[0]] = v[1]
    }
    // set of destination airports
    destset := map[string]bool{}
    // for each ticket
    for _, v := range tickets {
        // add destination airport to set
        destset[v[1]] = true
    }
    // holds itinerary
    it := [][]string{}
    // consider each departure and arrival pair
    for k := range tickmap {
        // source airport will not be in destination set
        if _, ok := destset[k]; ok {
            continue
        }
        // build up itinerary
        src := k
        for {
            // get destination airport
            dest, ok := tickmap[src]
            // if not ok
            if !ok {
                // done with itinerary
                break
            }
            // add source -> destination
            it = append(it, []string{src, dest})
            // set source airport to destination
            src = dest
        }
    }
    // return itinerary
    return it
}

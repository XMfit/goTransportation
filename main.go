/*
 * Author: Thomas Johnson, thomasjohnso2020@my.fit.edu
 * Course: CSE 4250, Fall 2023
 * Project: Project #2, Manatee Transportation
 * Implementation: go1.18.1 linux/amd64
 * https://en.wikipedia.org/wiki/Knapsack_problem helped me with the knapsack problem
 * https://www.geeksforgeeks.org/bin-packing-problem-minimize-number-of-used-bins/# gave me an idea of how to solve bin packing
 * https://sites.radford.edu/~nokie/classes/360/dp-memoized.html helped me understand memoization, and dp
 * https://www.interviewcake.com/concept/java/memoization basically rewording of the above
 */

package main

import (
	"fmt"
)

// Recursive function to determine the optimal arrangement of manatees in the boat
func loadManatees(tubsLength int, manatees []int, portLength int, starboardLength int, index int, memo map[string][]interface{}) (int, []string) {
	// if all manatees are considered, return 0 and an empty arrangement
	if index == len(manatees) {
		return 0, []string{}
	}

	// Create a unique key for the current state (manatee index, port length, starboard length)
	key := fmt.Sprintf("%d_%d_%d", index, portLength, starboardLength)

	// If the current state has been computed before, return the stored result from the memo
	if result, found := memo[key]; found {
		return result[0].(int), result[1].([]string)
	}

	// Init portside as 0, as we have 0 manatees on the port side
	// Create a slice of strings to store manatees on the port side
	portSide, portArrangement := 0, []string{}
	if portLength+manatees[index] <= tubsLength { // Check if the manatee can fit on the port side
		// make a recursive call to consider the next managee, with the updated length, and portSide
		portSide, portArrangement = loadManatees(tubsLength, manatees, portLength+manatees[index], starboardLength, index+1, memo)
		portSide += 1 // inc portSide
		// Update the arrangement with the current manatee loaded to the port.
		//portArrangement = append([]string{fmt.Sprintf("%d port", manatees[index])}, portArrangement...)
		portArrangement = append([]string{fmt.Sprintf("port")}, portArrangement...)
	}

	// Same process but for the starboard side
	starboardSide, starboardArrangement := 0, []string{}
	if starboardLength+manatees[index] <= tubsLength {

		starboardSide, starboardArrangement = loadManatees(tubsLength, manatees, portLength, starboardLength+manatees[index], index+1, memo)
		starboardSide += 1

		//starboardArrangement = append([]string{fmt.Sprintf("%d starboard", manatees[index])}, starboardArrangement...)
		starboardArrangement = append([]string{fmt.Sprintf("starboard")}, starboardArrangement...)
	}

	// Compare the port and starboard side arrangements and choose the one that allows more manatees to be loaded
	// If both are equal just do portSide
	if portSide >= starboardSide {
		memo[key] = []interface{}{portSide, portArrangement} // Memoize the result for the current state
		return portSide, portArrangement
	} else {
		memo[key] = []interface{}{starboardSide, starboardArrangement}
		return starboardSide, starboardArrangement
	}
}

func main() {

	// Read tub length input, and convert to cm
	var tubsLength int
	fmt.Scan(&tubsLength)
	tubsLength *= 100

	var manatees []int // init slice to store manatees

	var manateeLength int
	for {
		fmt.Scan(&manateeLength) // Continuously read manatee lengths until 0
		if manateeLength == 0 {
			break
		}
		manatees = append(manatees, manateeLength) // Add the manatee length to the slice
	}

	// Initialize a map for memoization, and specify the value as an interface
	// An interface{} allows for multiple values of unspecified type to be stored
	memo := make(map[string][]interface{})

	// Call the function to get the optimal arrangement
	totalManatees, arrangement := loadManatees(tubsLength, manatees, 0, 0, 0, memo)

	fmt.Println(totalManatees)
	for _, side := range arrangement {
		fmt.Println(side) // Print the arrangement for each manatee
	}
}

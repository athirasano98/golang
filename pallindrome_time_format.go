//Given a string A which represents a time in 24 hour HH:MM format.
// Find the minimum number of minutes need to pass to reach palindromic time.
// Input Format
// Given a string A.

// Output Format
// Return an integer.

// Example Input
// Input 1:
// A = "23:59"
// Input 2:

// A = "21:00"

// Example Output
// Output 1:
// 1
// Output 2:

// 12

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(s int) int {
	sum := 0
	for {
		d := s % 10
		sum = (sum * 10) + d
		s = s / 10
		if s == 0 {
			return sum
		}
	}

}
func main() {
	var a = "20:59"
	val := strings.Split(a, ":")
	intSlice := make([]int, len(val))
	for i, s := range val {
		intSlice[i], _ = strconv.Atoi(s)
	}
	stringRev := reverse(intSlice[0])
	fmt.Println(stringRev)
	if stringRev == intSlice[1] {
		fmt.Println("already in pallindrome")
	} else {
		if intSlice[1] > stringRev {
			h := intSlice[0] + 1
			if h == 24 {
				h = 0
			}
			diff := (60 - intSlice[1]) + reverse(h)
			fmt.Println(diff, "minutes is required to become pallindrome")

		} else {
			diff := (stringRev - intSlice[1])
			fmt.Println(diff, "minutes is required to become pallindrome")
		}
	}

}

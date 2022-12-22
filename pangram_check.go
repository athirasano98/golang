// Given a sentence represented as an array A of
// strings that contains all lowercase alphabets.
// Chech if it is a pangram or not.
// A pangram is a unique sentence in which every
// letter of the lowercase alphabet is used at least once.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var A = []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	s := strings.Join(A, "")
	for i := 'a'; i <= 'z'; i++ {
		str := string(i)
		if !strings.Contains(s, str) {
			fmt.Println(str)
			fmt.Println("invalid")
			return
		}

	}
	fmt.Println("valid")

}

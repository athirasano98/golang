package main

// import (
// 	"fmt"
// )

// type Stack struct {
// 	items []string
// }

// var matchingPair = map[string]string{
// 	"{": "}",
// 	"[": "]",
// 	"(": ")",
// }

// func (stack *Stack) Peek() string {

// 	if len(stack.items) == 0 {
// 		return " "
// 	}
// 	return stack.items[len(stack.items)-1]
// }
// func (stack *Stack) Push(item string) {

// 	stack.items = append(stack.items, item)
// }

// func (stack *Stack) Pop() string {

// 	if len(stack.items) == 0 {
// 		return " "
// 	}
// 	lastItem := stack.items[len(stack.items)-1]
// 	stack.items = stack.items[:len(stack.items)-1]
// 	return lastItem
// }
// func mapValues(pairs map[string]string, val string) bool {
// 	for _, value := range pairs {
// 		if val == value {
// 			return true
// 		}

// 	}
// 	return false
// }
// func main() {
// 	var (
// 		stack Stack
// 		value string
// 	)
// 	fmt.Println("enter the string")
// 	_, err := fmt.Scan(&value)

// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, ch := range value {
// 		val := string(ch)
// 		if _, ok := matchingPair[val]; ok {

// 			stack.Push(val)

// 		} else if ok := mapValues(matchingPair, val); ok {

// 			char := stack.Peek()

// 			if char == " " || val != matchingPair[char] {
// 				fmt.Println("string is not valid")
// 				return
// 			}

// 			stack.Pop()

// 		}
// 	}

// 	if len(stack.items) != 0 {
// 		fmt.Println("string is not  valid")
// 	} else {
// 		fmt.Println("string is  valid")
// 	}

// }

package main

import (
	"fmt"
	"reflect"
	"sort"
)

type KeyValue struct {
	Key   string
	Value int
}

func main() {
	var values = map[string]int{
		"apple":     1,
		"mango":     1,
		"orange":    10,
		"grapes":    1,
		"pineapple": 1,
		"guava":     8,
	}

	s := make([]KeyValue, 0, len(values))

	for k, v := range values {
		s = append(s, KeyValue{k, v})
	}

	swapSlice := reflect.Swapper(s)

	sort.Slice(s, func(i, j int) bool {
		if s[i].Value == s[j].Value {
			if s[i].Key > s[j].Key {
				swapSlice(i, j)
			}

		}
		return s[i].Value >= s[j].Value
	})

	for _, v := range s {
		fmt.Println(v.Key, "-", v.Value)
	}
}

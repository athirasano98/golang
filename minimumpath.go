package main

import "fmt"

func minimumPath(total int, start int, stop int, cost []int) int {
	sum1 := 0
	sum2 := 0
	for i := start; i < stop; i++ {
		sum1 += cost[i-1]
	}
	sum2 = total - sum1
	if sum1 > sum2 {
		return sum2
	}
	return sum1
}

func main() {
	var (
		n         int
		start     int
		stop      int
		totalCost int
	)
	fmt.Println("Enter the number of stations")
	fmt.Scan(&n)
	fmt.Println("enter the start station")
	fmt.Scan(&start)
	fmt.Println("enter the stop station")
	fmt.Scan(&stop)
	fmt.Println("enter the cost")
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cost[i])
		totalCost += cost[i]
	}
	fmt.Println(totalCost)
	minCost := minimumPath(totalCost, start, stop, cost)
	fmt.Println("minimum cost from", start, "to", stop, "is", minCost)
}

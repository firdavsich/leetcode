package main

import "fmt"

func maxProfit(prices []int) int {
	var price int
	maxProf := 0
	currMin := prices[0]

	for i := 1; i < len(prices); i++ {
		price = prices[i]

		if curProf := price - currMin; curProf > maxProf {
			maxProf = curProf
		}
		if price < currMin {
			currMin = price
		}
	}

	return maxProf
}

func main() {
	fmt.Printf("Result: %d, Expect: %d\n", maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	fmt.Printf("Result: %d, Expect: %d\n", maxProfit([]int{7, 6, 4, 3, 1}), 0)
	fmt.Printf("Result: %d, Expect: %d\n", maxProfit([]int{2, 1, 4}), 3)
	fmt.Printf("Result: %d, Expect: %d\n", maxProfit([]int{1}), 0)

}

package main

import "fmt"

func reverse(x int) int {
	var pop, rev int
	if x > 0 {
		for x > 0 {
			pop = x % 10
			x /= 10
			rev = rev*10 + pop
		}
	} else {
		for x < 0 {
			pop = x % 10
			x /= 10
			rev = rev*10 + pop
		}
	}

	if -2147483648 > rev || rev > 2147483647 {
		return 0
	}
	return rev
}

func main() {
	fmt.Println(reverse(123), reverse(123) == 321)
	fmt.Println(reverse(-123), reverse(-123) == -321)
	fmt.Println(reverse(120), reverse(120) == 21)
	fmt.Println(reverse(0), reverse(0) == 0)
	fmt.Println(reverse(1534236469), reverse(1534236469) == 0)

}

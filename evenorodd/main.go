package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5 ,6, 7, 8, 9, 10}
	evenOrOdd(nums)
}

func evenOrOdd(nums []int) {
	for _, n := range nums {
		if n % 2 == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}
}
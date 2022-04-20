package main

import "fmt"

func findNumbers(nums []int) int {
	counter := 0
	digitCouter := 1
	res := 0
	for _, v := range nums {
		if v < 10 {
			continue
		}
		res = v
		for {
			// fmt.Println(res)
			res = res / 10
			digitCouter++
			if res < 10 {
				break
			}
		}
		// fmt.Println(digitCouter)

		if digitCouter%2 == 0 {
			counter++
		}
		digitCouter = 1
	}
	return counter
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	result := findNumbers(nums)
	fmt.Printf("even nums in array: %d\n", result)
}

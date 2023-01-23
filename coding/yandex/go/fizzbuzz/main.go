package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Println("FizzBuzz")
		case 0 == i%3:
			fmt.Println("Fizz")
		case 0 == i%5:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

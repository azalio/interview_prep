package main

import "fmt"

func reverse(s []byte) {
	if len(s) <= 1 {
		return
	}

	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	reverse(s[1 : len(s)-1])

}

func main() {
	str := []byte("hello")
	reverse(str)
	fmt.Println(string(str))
}

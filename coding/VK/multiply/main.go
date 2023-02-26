package multiply

import "fmt"

func printMultTable(n int) string {

	var result string
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			result += fmt.Sprintf("%d\t", i*j)
		}
		result += fmt.Sprintln()
	}

	return result
}

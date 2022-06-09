package pascalTriangle

func generate(numRows int) [][]int {

	res := make([][]int, numRows)

	for x := 0; x < numRows; x++ {
		res[x] = make([]int, x+1)
		for y := 0; y <= x; y++ {
			if y == 0 || y == x {
				res[x][y] = 1
			} else {
				res[x][y] = res[x-1][y-1] + res[x-1][y]
			}
		}
	}

	return res
}

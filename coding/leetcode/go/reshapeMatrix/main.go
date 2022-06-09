package reshapeMatrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n := len(mat[0])
	var mm, nn int

	var lenOfMatrix int

	lenOfMatrix = len(mat) * len(mat[0])

	if r*c != lenOfMatrix {
		return mat
	}

	result := make([][]int, r)

	for rr := 0; rr < r; rr++ {
		for cc := 0; cc < c; cc++ {

			y := mat[mm][nn]
			result[rr] = append(result[rr], y)
			nn++
			if nn >= n {
				nn = 0
				mm++
			}
		}
	}

	return result
}

func matrixReshape2(mat [][]int, r int, c int) [][]int {
	n := len(mat)
	m := len(mat[0])
	if r*c != m*n {
		return mat
	}

	res := make([][]int, r)

	for x := 0; x < r; x++ {
		res[x] = make([]int, c)
	}

	for i := 0; i < r*c; i++ {
		res[i/c][i%c] = mat[i/m][i%m]
	}
	return res

}

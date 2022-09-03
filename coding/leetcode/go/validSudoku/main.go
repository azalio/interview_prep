package validSudoku

import (
	"errors"
)

func getNum(s byte) (int, error) {
	if s == 46 {
		return -1, errors.New("not digit")
	} else {
		return int(s), nil
	}
}

func isValidSudoku(board [][]byte) bool {
	sudoku := make(map[int]bool)

	for _, h := range board {
		for _, z := range h {
			n, err := getNum(z)
			if err != nil {
				continue
			} else {
				_, ok := sudoku[n]
				if ok {
					return false
				}
				sudoku[n] = true
			}
		}
		sudoku = make(map[int]bool)
	}

	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n, err := getNum(board[i][j])
			if err != nil {
				continue
			} else {
				_, ok := sudoku[n]
				if ok {
					return false
				}
				sudoku[n] = true
			}
		}

		sudoku = make(map[int]bool)
	}

	for i := 0; i < 9; i = i + 3 {
		nineToTheree := board[i : i+3]

		for j := 0; j < 9; j = j + 3 {
			for _, h := range nineToTheree {
				txt := h[j : j+3]
				for _, n := range txt {
					n, err := getNum(n)
					if err != nil {
						continue
					} else {
						_, ok := sudoku[n]
						if ok {
							return false
						}
						sudoku[n] = true
					}
				}
			}

			sudoku = make(map[int]bool)

		}

	}

	return true
}

//func isValidSudoku2(board [][]byte) bool {
//
//	seen := make(map([]))
//
//	return true
//}

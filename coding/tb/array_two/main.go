package array_two

//а вторая [-5, -3, 0, 1, 4], массив сортированый, надо возвести их в квадрат
// и сохранить сортированность,
//так как отрицательные числа станут положительными
//ну и тоже с лишней памятью/без лишней памяти

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(s []int) []int {
	rIdx := 0
	lIdx := len(s) - 1
	result := make([]int, len(s), len(s))

	for i := len(s) - 1; rIdx <= lIdx; i-- {
		if Abs(s[rIdx]) >= Abs(s[lIdx]) {
			result[i] = s[rIdx] * s[rIdx]
			rIdx++
		} else {
			result[i] = s[lIdx] * s[lIdx]
			lIdx--
		}
	}
	return result
}

package array_one

import (
	"sort"
)

//дан массив чисел, [4,3,0,1] несортированый, от 0 до N надо найти число которого не хватает,
//ну и там по скорости или по времени лучшее решение предложить
//ну и пограничные случаи типа 0 и N учесть

func findNumber(s []int) int { // O(n) - cpu, O(n) - mem
	sM := make(map[int]bool)
	for _, v := range s {
		sM[v] = true
	}

	for i := 0; i < len(s); i++ {
		_, ok := sM[i]
		if !ok {
			return i
		}
	}
	return 0
}

func findNumber2(s []int) int { // O(nlogn) - cpu, O(log n) - mem
	sort.Ints(s)
	ii := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ii {
			return i
		}
		ii++
	}
	return 0
}

func findNumber3(s []int) int { // O(N) - cpu, O(1) - mem

	if len(s) == 1 {
		return 0
	}

	sumOfWholeSlice := (len(s)) * (len(s) + 1) / 2
	//fmt.Println(sumOfWholeSlice)
	var sumOfRealSlice int
	for _, v := range s {
		sumOfRealSlice += v
	}

	//fmt.Println(sumOfRealSlice)
	return sumOfWholeSlice - sumOfRealSlice
}

package day_01

func SolvePart2() int64 {
	var ans int64
	var countB = make(map[int64]int64)

	listA, listB := ParseInput()

	for _, val := range listB {
		countB[val] += 1
	}

	for _, val := range listA {
		ans += val * countB[val]
	}

	return ans
}

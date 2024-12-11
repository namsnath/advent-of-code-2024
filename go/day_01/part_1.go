package day_01

import (
	"math"
	"slices"
)

func SolvePart1() int64 {
	var ans int64

	listA, listB := ParseInput()
	slices.Sort(listA)
	slices.Sort(listB)

	for idx := range len(listA) {
		ans += int64(math.Abs(float64(listA[idx] - listB[idx])))
		// fmt.Printf("%d => A=%d, B=%d => ans=%d\n", idx, listA[idx], listB[idx], ans)
	}

	return ans
}

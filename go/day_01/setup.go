package day_01

import (
	"advent-of-code-2024/util"
	_ "embed"
	"errors"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputContent string

func ParseInput() (listA, listB []int64) {
	lines := strings.Split(inputContent, "\n")
	listA = make([]int64, len(lines))
	listB = make([]int64, len(lines))

	for idx, line := range lines {
		values := strings.Split(line, "   ")
		if len(values) != 2 {
			panic(errors.New("Invalid Line: " + line))
		}

		listA[idx] = util.Must(strconv.ParseInt(values[0], 10, 64))
		listB[idx] = util.Must(strconv.ParseInt(values[1], 10, 64))
	}

	return listA, listB
}

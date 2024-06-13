// https://school.programmers.co.kr/learn/courses/30/lessons/42746

import (
	"sort"
	"strconv"
)

func solution(numbers []int) string {
	result := ""

	sort.Slice(numbers, func(i, j int) bool {
		a := strconv.Itoa(numbers[i])
		b := strconv.Itoa(numbers[j])

		return a+b > b+a
	})

	if numbers[0] == 0 {
		return "0"
	}
	for _, n := range numbers {
		result += strconv.Itoa(n)
	}

	return result
}
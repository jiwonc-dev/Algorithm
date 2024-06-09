// https://school.programmers.co.kr/learn/courses/30/lessons/42748

import "sort"

/*
	input : array, commands
	output : result array
	O(N^2)
*/
func solution(array []int, commands [][]int) []int {
	result := []int{}

	for _, cmd := range commands {
		i, j, k := cmd[0], cmd[1], cmd[2]
		// shallow copy
		parsedArr := append([]int{}, array[(i-1):j]...)

		sort.Ints(parsedArr)
		result = append(result, parsedArr[k-1])
	}
	return result
}

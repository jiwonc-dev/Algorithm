// https://school.programmers.co.kr/learn/courses/30/lessons/42747
import "sort"

func solution(citations []int) int {
	count := len(citations)

	sort.Ints(citations)

	for i := 0; i < count; i++ {
		if citations[i] == 0 {
			continue
		}
		restCount := count - i
		if restCount <= citations[i] {
			return restCount
		}
	}

	return 0
}
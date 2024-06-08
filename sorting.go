package main

import (
	"log"
	"strconv"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 8, 2, 6, 4}

	// selectionSort(&arr)
	// bubbleSort(&arr)
	// insertionSort(&arr)
	quickSort(&arr, 0, len(arr)-1)

	printArr(arr)
}

/*
input : array
output : void (print array)
time : O(N)
`
*/
func printArr(arr []int) {
	result := "{ "
	for _, n := range arr {
		result += strconv.Itoa(n) + " "
	}
	result += "}"
	log.Println(result)
}

/*
input : 2 int pointers
output : void
time : O(1)
*/
func swap(val1, val2 *int) {
	temp := *val1
	*val1 = *val2
	*val2 = temp
}

/*
input : unsorted array pointer
output : void (just sorting)
time : O(N^2)
*/
func selectionSort(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n; i++ {
		min := (*arr)[i]
		// find the smallest one
		for j := i + 1; j < n; j++ {
			if min > (*arr)[j] {
				swap(&min, &(*arr)[j])
			}
		}
		(*arr)[i] = min
	}
}

/*
input : unsorted array pointer
output : void (just sorting)
time : O(N^2)
*/
func bubbleSort(arr *[]int) {
	n := len(*arr)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			left, right := (*arr)[j], (*arr)[j+1]
			if left > right {
				swap(&(*arr)[j], &(*arr)[j+1])
			}
		}
	}
}

/*
input : unsorted array pointer
output : void (just sorting)
time : O(N^2)
*/
func insertionSort(arr *[]int) {
	n := len(*arr)
	for i := 1; i < n; i++ {
		curIdx := i
		curVal := (*arr)[curIdx]
		for j := i - 1; j >= 0; j-- {
			left := (*arr)[j]
			if left > curVal {
				swap(&(*arr)[curIdx], &(*arr)[j])
				curIdx = j
			} else {
				break
			}
		}
	}
}

/*
input : unsorted array pointer, start index, end index
output : void (just sorting)
time : O(NlogN)
*/
func quickSort(arr *[]int, start, end int) {
	// if the array has just one value, return
	if start >= end {
		return
	}
	pivot := start
	left, right := start+1, end

	for left < right {
		// left pointer finds the bigger value than pivot value
		for left <= end && (*arr)[left] < (*arr)[pivot] {
			left++
		}
		// right pointer finds the smaller value than pivot value
		for right > start && (*arr)[right] > (*arr)[pivot] {
			right--
		}

		// if the pointers are crossed, switch the pivot
		if left > right {
			swap(&(*arr)[pivot], &(*arr)[right])
		} else {
			swap(&(*arr)[left], &(*arr)[right])
		}
	}

	quickSort(arr, start, right-1)
	quickSort(arr, right+1, end)
}

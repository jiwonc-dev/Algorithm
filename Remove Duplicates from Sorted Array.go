func removeDuplicates(nums []int) int {
	numOfElements := len(nums)
	if numOfElements == 1 {
		return 1
	}

	numOfDublicates := 0
	newNums := []int{}
	numOfUnique := 0

	for left := 0; left < numOfElements; {
		leftNum := nums[left]
		newNums = append(newNums, leftNum)
		numOfUnique++

		right := left + 1
		for right < numOfElements && nums[right] == leftNum {
			right++
			numOfDublicates++
		}

		left = right
	}

	// change nums
	for i := 0; i < numOfElements; i++ {
		if i >= numOfUnique {
			break
		}
		nums[i] = newNums[i]
	}

	return numOfUnique
}
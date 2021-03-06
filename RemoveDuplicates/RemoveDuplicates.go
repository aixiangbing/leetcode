package RemoveDuplicates

func removeDuplicates(nums []int) int  {

	if len(nums) == 0 {
		return 0
	}

	length := len(nums)

	lastNum := nums[length - 1]

	i := 0

	for i = 0; i < length -1; i++ {
		if nums[i] == lastNum {
			break
		}
		// if next element equals current , remove the next element
		if nums[i + 1] == nums[i] {
			removeElements(nums, i+ 1, nums[i])
		}
	}

	return i + 1
}

func removeElements(nums []int, start int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	j := start
	for i := start; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		} else {
			j++
		}
	}
	return j
}

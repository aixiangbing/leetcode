package ThreeNumberSum


import (
	"sort"
)

func threeSum(nums[]int) [][]int  {

	sort.Ints(nums)

	ans := make([][]int, 0)

	for first := 0; first < len(nums) - 2; first++ {

		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}

		target := 0 - nums[first]

		third := len(nums) - 1

		for second := first + 1; second < len(nums) - 1; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}

			for  second < third && nums[second] + nums[third] > target  {
				third--
			}

			if second == third {
				break
			}

			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
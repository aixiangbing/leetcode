package maximumGap

import "sort"

func maximumGap(nums []int) int {
	sort.Ints(nums)

	ans := 0
	for i:=0; i < len(nums) - 1; i++ {
		res := nums[i + 1] - nums[i]
		if res > ans {
			ans = res
		}
	}
	return ans
}


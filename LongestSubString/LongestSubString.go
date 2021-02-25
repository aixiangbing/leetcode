package LongestSubString

func lengthOfLongestSubstring(s string) int  {

	allString := make(map[string]int)

	//定义左指针
	i := -1

	//定义返回结果
	res := 0

	for j := 0; j < len(s); j++ {
		if _, ok := allString[string(s[j])]; ok {
			i = max(i, allString[string(s[j])])
		}
		allString[string(s[j])] = j

		res = max(res, j - i)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

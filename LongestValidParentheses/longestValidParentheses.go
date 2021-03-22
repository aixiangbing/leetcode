package LongestValidParentheses

/**

给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

 

示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses

 */
func longestValidParentheses(s string) int {

	maxAns := 0

	stack := []int{}

	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack) - 1]

			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i - stack[len(stack) - 1])
			}
		}
	}
	return maxAns
}

func max(x, y int) int  {
	if x > y {
		return x
	}
	return y
}
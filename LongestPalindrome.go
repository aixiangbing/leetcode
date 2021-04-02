package leetcode

// 动态规划
func longestPalindrome(s string) string  {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i + l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l + 1 > len(ans) {
				ans = s[i:i+l+1]
			}
		}
	}
	return ans

}


// 中心回朔
func longestPalindromeTwo(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i + 1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
	return left + 1, right - 1
}


//  暴力解决

func longestPalindromeThree(s string)  string {

	// 定义字符串的长度
	len := len(s)
	if len < 2 {
		return s
	}

	maxLen := 1
	begin := 0

	for i := 0; i < len - 1; i++ {
		for j := i + 1; j < len; j++ {
			if j - i + 1 > maxLen && validPalindromic(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin: begin+maxLen]
}

func validPalindromic(s string, left int, right int ) bool  {

	for left < right {

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}





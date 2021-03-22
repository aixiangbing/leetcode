package main

import (
	"fmt"
	//"leetcode/leetcode/RandPrint"
)

func main()  {
	//p := linkedList.MyLinkedList{}
	//p.AddAtHead(1)
	//p.AddAtTail(3)
	//p.AddAtIndex(1, 2)
	//p.Get(1)
	//p.DeleteAtIndex(1)
	//fmt.Println(p.Get(1))
	//
	//l := letterCombination.LetterCombinations{}
	//s := l.LetterCombination("234")
	//
	//fmt.Println(s)

	//nums := []int{2, 3, 4, 11, 12, 9, 6, 100, 99, 56, 1010, 321, 1, 19}
	//
	//selfSort := RadixSort.RadixSort{}
	//
	//res := selfSort.Sort(nums)
	//
	//fmt.Println(res)
	//
	//randPrint := RandPrint.RandPrint{}
	//randPrint.Print()

	//randNum := RandPrint.RandNum{}
	//
	//randNum.RandPrintNum()

	s := "2312312"
	fmt.Println(s[1:2])
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
	return s[begin: begin + maxLen]
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
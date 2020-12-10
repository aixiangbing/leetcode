package main

import (
	"fmt"
	"leetcode/leetcode/linkedList"
	"leetcode/leetcode/letterCombination"
)

func main()  {
	p := linkedList.MyLinkedList{}
	p.AddAtHead(1)
	p.AddAtTail(3)
	p.AddAtIndex(1, 2)
	p.Get(1)
	p.DeleteAtIndex(1)
	fmt.Println(p.Get(1))

	l := letterCombination.LetterCombinations{}
	s := l.LetterCombination("234")

	fmt.Println(s)



}

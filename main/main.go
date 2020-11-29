package main

import (
	"fmt"
	"leetcode/leetcode/linkedList"
)

func main()  {
	p := linkedList.MyLinkedList{}
	p.AddAtHead(1)
	p.AddAtTail(3)
	p.AddAtIndex(1, 2)
	p.Get(1)
	p.DeleteAtIndex(1)
	fmt.Println(p.Get(1))

}

package doubleLinkedList

import "math"

type Node struct {
	prev *Node
	next *Node
	val  int
}

type MyLinkedList struct {
	head *Node
	tail *Node
	len  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	head := new(Node)
	tail := new(Node)
	head.val = math.MinInt32
	head.next = tail
	tail.val = math.MaxInt32
	tail.prev = head
	return MyLinkedList{head, tail, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.len {
		return -1
	}

	curr := this.head.next
	ans := -1

	for i := index; i >= 0; i-- {
		if curr == nil {
			break
		}
		ans = curr.val
		curr = curr.next
	}

	return ans
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	next := this.head.next

	node := new(Node)
	node.val = val
	node.prev = this.head
	node.next = next

	this.head.next = node
	next.prev = node

	this.len++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	prev := this.tail.prev

	node := new(Node)
	node.val = val
	node.prev = prev
	node.next = this.tail

	this.tail.prev = node
	prev.next = node

	this.len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.len {
		this.AddAtTail(val)
		return
	}

	curr := this.head

	for i := 0; i <= index; i++ {
		curr = curr.next
	}

	prev := curr.prev

	node := new(Node)
	node.val = val
	node.prev = prev
	node.next = curr

	curr.prev = node
	prev.next = node

	this.len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.len || index < 0 {
		return
	}

	curr := this.head

	for i := 0; i <= index; i++ {
		curr = curr.next
	}

	prev := curr.prev
	next := curr.next

	prev.next = next
	next.prev = prev

	this.len--
}

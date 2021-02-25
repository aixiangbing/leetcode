package ValidateStackSequences

import "container/list"

func validateStackSequences(pushed []int, popped []int) bool  {
	stack := list.New()
	i := 0
	for j := 0; j < len(pushed); j++ {
		stack.PushBack(pushed[j])
		for stack.Len() > 0 && stack.Back().Value == popped[i] {
			stack.Remove(stack.Back())
			i++
		}
	}

	if stack.Len() == 0 {
		return true
	}
	return false
}

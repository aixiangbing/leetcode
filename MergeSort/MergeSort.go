package MergeSort

func mergeSort(arr []int) []int {

	arrLength := len(arr)

	if arrLength < 2 {
		return arr
	}

	middle := arrLength / 2

	left := arr[0:middle]

	right := arr[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge (left []int, right []int) []int {

	var result []int

	for len(left) != 0 && len(right) != 0{

	}
}

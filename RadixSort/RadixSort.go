package RadixSort

import "fmt"

type SelfSort interface {
	Sort(sortArray []int) []int
}

type RadixSort struct {
}

func (this RadixSort) Sort(sortArray []int) []int {
	arr := sortArray
	maxDigit := this.getMaxDigit(arr)
	fmt.Println(maxDigit)
	return this.radixSort(arr, maxDigit)
}

func (this RadixSort) getMaxDigit(arr []int) int {
	maxValue := this.getMaxValue(arr)
	maxDigit := this.getMaxValueLength(maxValue)

	return maxDigit
}

func (this RadixSort) getMaxValue(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxValue < arr[i] {
			maxValue = arr[i]
		}
	}
	return maxValue
}

func (this RadixSort) getMaxValueLength(maxValue int) int {
	if maxValue == 0 {
		return 1
	}

	maxValueLength := 0

	for i := maxValue; i != 0; i /= 10 {
		maxValueLength++
	}
	return maxValueLength
}

func (this RadixSort) radixSort(arr []int, maxDigit int) []int {

	mod := 10
	dev := 1

	for i := 0; i < maxDigit; i++ {
		counter := make([][]int, len(arr) + 10)
		for _, value := range arr {
			bucket := (value%mod)/dev + 10
			counter[bucket] = append(counter[bucket], value)
		}

		pos := 0
		for _, radixArr := range counter {
			for _, value := range radixArr {
				arr[pos] = value
				pos++
			}
		}
		mod *= 10
		dev *= 10
	}
	return arr
}

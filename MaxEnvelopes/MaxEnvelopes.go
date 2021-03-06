package MaxEnvelopes

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sums := make([]int, 0)
	result := make(map[string]int, 0)
	for i := 0 ; i < len(envelopes); i++ {
		temp := 0
		for j := 0; j < len(envelopes[i]); j++ {
			temp += envelopes[i][j]
		}
		sums = append(sums, temp)
	}

	sort.Ints(sums)
}

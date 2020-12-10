package letterCombination

import "fmt"

type LetterCombinations struct {
	letters map[string] []string
	combinations []string
}

func Constructor() LetterCombinations {
	return LetterCombinations{letters: make(map[string] []string, 8), combinations: []string{}}
}

func (this *LetterCombinations) LetterCombination(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	this.letters = make(map[string] []string, 8)
	this.combinations = make([]string, 0)

	this.letters["2"] = append(this.letters["2"], "a", "b", "c")
	this.letters["3"] = append(this.letters["3"], "d", "e", "f")
	this.letters["4"] = append(this.letters["4"], "g", "h", "i")
	this.letters["5"] = append(this.letters["5"], "j", "k", "l")
	this.letters["6"] = append(this.letters["6"], "m", "n", "o")
	this.letters["7"] = append(this.letters["7"], "p", "q", "r", "s")
	this.letters["8"] = append(this.letters["8"], "t", "u", "v")
	this.letters["9"] = append(this.letters["9"], "w", "x", "y", "z")

	this.GeneratorLetterCombination(digits, 0, "")
	fmt.Println("------")
	fmt.Println(digits[0])
	fmt.Println("-------")
	return this.combinations
}

func (this *LetterCombinations) GeneratorLetterCombination(digits string, index int, combination string)  {

	if index == len(digits) {
		this.combinations = append(this.combinations, combination)
	} else {
		digit := string(digits[index])
		letter := this.letters[digit]
		for i := 0; i < len(letter); i++ {
			this.GeneratorLetterCombination(digits, index + 1, combination + string(letter[i]))
		}
	}
}
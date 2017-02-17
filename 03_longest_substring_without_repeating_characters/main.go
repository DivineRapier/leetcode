package main

import "fmt"

func main() {
	// fmt.Println(LongestSubstringWithoutRepeatingCharacters("aaabbbbb"))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("abca"))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("abcabc"))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("abcabcdefg"))

}

func LongestSubstringWithoutRepeatingCharacters(s string) (int, int) {

	const MaxCharactersCount = 256

	var listOfLastAppearPosition [MaxCharactersCount]int

	setDefaultPositiong(listOfLastAppearPosition[:])

	return longestSubstringWithoutRepeatingCharacters([]byte(s), listOfLastAppearPosition[:])
}

// longestSubstringWithoutRepeatingCharacters
// bs : 字符串
// posList : 字符最后出现的位置
func longestSubstringWithoutRepeatingCharacters(bs []byte, posList []int) (int, int) {

	startMaxPos, endMaxPos := 0, 0

	lastRepeatPos := -1

	for index, value := range bs {

		lastPos := posList[int(value)]

		if lastPos != -1 && lastRepeatPos < lastPos {
			lastRepeatPos = lastPos
		}

		if index-lastRepeatPos > endMaxPos-startMaxPos {
			endMaxPos, startMaxPos = index, lastRepeatPos
		}

		posList[value] = index
	}

	return startMaxPos + 1, endMaxPos + 1
}

func setDefaultPositiong(slice []int) {
	for i := range slice {
		slice[i] = -1
	}
}

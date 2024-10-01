package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var punctuationMarks = regexp.MustCompile(`[[:punct:]]`)

func Top10(wordSequence string) []string {
	var result []string

	wordSequence = strings.ToLower(wordSequence)
	wordArray := strings.Fields(wordSequence)

	wordMap := make(map[string]int)

	for _, word := range wordArray {
		sequenceRunes := []rune(word)

		transformedWord := proceedCharacter(sequenceRunes)

		if len(transformedWord) != 0 {
			if _, ok := wordMap[transformedWord]; ok {
				wordMap[transformedWord]++
			} else {
				wordMap[transformedWord] = 1
				result = append(result, transformedWord)
			}
		}
	}

	result = customSort(result, wordMap)

	if len(result) > 10 {
		return result[:10]
	}

	return result
}

func proceedCharacter(r []rune) string {
	transformedWord := ""

	for i := 0; i < len(r); i++ {
		if punctuationMarks.MatchString(string(r[i])) {
			if len(r) != 1 && !((i == 0 && r[i] != r[i+1]) || (i == len(r)-1 && r[i] != r[i-1])) {
				transformedWord += string(r[i])
			}
		} else {
			transformedWord += string(r[i])
		}
	}

	return transformedWord
}

func customSort(result []string, wordMap map[string]int) []string {
	sort.Slice(result, func(i int, j int) bool {
		if wordMap[result[i]] > wordMap[result[j]] {
			return true
		} else if wordMap[result[i]] == wordMap[result[j]] {
			if result[i] < result[j] {
				return true
			}
		}
		return false
	})

	return result
}

package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var punctuationMark = regexp.MustCompile(`[[:punct:]]`)

var punctuationMarks = regexp.MustCompile(`^[[:punct:]]{2,}$`)

func Top10(wordSequence string) []string {
	var result []string

	wordSequence = strings.ToLower(wordSequence)
	wordArray := strings.Fields(wordSequence)

	wordMap := make(map[string]int)

	for _, word := range wordArray {
		transformedWord := convertWord(word)

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

func convertWord(word string) string {
	if punctuationMarks.MatchString(word) {
		return word
	}

	return strings.TrimFunc(word, func(r rune) bool {
		return punctuationMark.MatchString(string(r))
	})
}

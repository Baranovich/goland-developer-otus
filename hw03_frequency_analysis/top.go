package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[a-zа-яё]+[^a-zа-яё]*[a-zа-яё]+[\-a-zа-яё]*|\-{2,}|[a-zа-яё]`)

func Top10(wordSequence string) []string {
	var result []string

	wordSequence = strings.ToLower(wordSequence)
	wordArray := strings.Fields(wordSequence)

	wordMap := make(map[string]int)

	for _, word := range wordArray {
		transformedWordArray := re.FindAllString(word, -1)

		if len(transformedWordArray) != 0 {
			transformedWord := transformedWordArray[0]

			if _, ok := wordMap[transformedWord]; ok {
				wordMap[transformedWord]++
			} else {
				wordMap[transformedWord] = 1
				result = append(result, transformedWord)
			}
		}
	}

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

	if len(result) > 10 {
		return result[:10]
	}

	return result
}

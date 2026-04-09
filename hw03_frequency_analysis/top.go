package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)
	resultMap := make(map[string]int)
	for _, word := range words {
		resultMap[word]++
	}

	type wordStatistic struct {
		word  string
		count int
	}

	wordStatistics := make([]wordStatistic, 0, len(resultMap))
	for word, count := range resultMap {
		wordStatistics = append(wordStatistics, wordStatistic{word, count})
	}

	sort.Slice(wordStatistics, func(i, j int) bool {
		if wordStatistics[i].count == wordStatistics[j].count {
			return wordStatistics[i].word < wordStatistics[j].word
		}
		return wordStatistics[i].count > wordStatistics[j].count
	})

	length := 10
	if len(wordStatistics) < length {
		length = len(wordStatistics)
	}

	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = wordStatistics[i].word
	}

	return result
}

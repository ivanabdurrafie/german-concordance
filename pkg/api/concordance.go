package api

import (
	"regexp"
	"sort"
	"strings"

	"slices"

	snowball "github.com/AlasdairF/Stemmer"
)

type ConcordanceRequest struct {
	Text string `json:"text" binding:"required"`
}

type WordInfo struct {
	Count      int    `json:"count"`
	Stem       string `json:"stem"`
	LineCounts []int  `json:"line_numbers"`
}

func Concordance(text string) map[string]WordInfo {
	concordance := make(map[string]WordInfo)
	lines := strings.Split(text, "\n")
	wordRegex := regexp.MustCompile(`[\p{L}]+`)

	stemmer, err := snowball.New("german")
	if err != nil {
		panic(err)
	}

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		words := wordRegex.FindAllString(line, -1)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			stemmed := stemmer.Stem(lowerWord)

			if info, exists := concordance[lowerWord]; exists {
				info.Count++
				if !contains(info.LineCounts, lineNum+1) {
					info.LineCounts = append(info.LineCounts, lineNum+1)
				}
				concordance[lowerWord] = info
			} else {
				concordance[lowerWord] = WordInfo{
					Count:      1,
					Stem:       stemmed,
					LineCounts: []int{lineNum + 1},
				}
			}
		}
	}

	for word, info := range concordance {
		sort.Ints(info.LineCounts)
		concordance[word] = info
	}
	return concordance
}

func contains(slice []int, value int) bool {
	return slices.Contains(slice, value)
}

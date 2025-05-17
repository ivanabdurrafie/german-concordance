package api

import (
	"net/http"
	"regexp"
	"sort"
	"strings"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/kljensen/snowball"
)

type ConcordanceRequest struct {
	Text string `json:"text" binding:"required"`
}

type WordInfo struct {
	Count      int    `json:"count"`
	Stem       string `json:"stem"`
	LineCounts []int  `json:"line_numbers"`
}

func ConcordanceHandler(c *gin.Context) {
	var req ConcordanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	concordance := Concordance(req.Text)
	c.JSON(http.StatusOK, concordance)
}

func Concordance(text string) map[string]WordInfo {
	concordance := make(map[string]WordInfo)
	lines := strings.Split(text, "\n")
	wordRegex := regexp.MustCompile(`[\p{L}]+`)

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		words := wordRegex.FindAllString(line, -1)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			stemmed, _ := snowball.Stem(lowerWord, "german", true)

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

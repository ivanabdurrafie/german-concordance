package api

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Health check
// @Description Check if the API is running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// ConcordanceHandler godoc
// @Summary Analyze German text
// @Description Generate a concordance for German text input
// @Tags concordance
// @Accept text/plain
// @Produce json
// @RequestBody string true "German text"
// @Success 200 {object} map[string]WordInfo
// @Failure 400 {object} map[string]string
// @Router /concordance [post]
func ConcordanceHandler(c *gin.Context) {
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	text := strings.TrimSpace(string(rawBody))
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty text input"})
		return
	}

	// Optional: wrap into struct if you want
	req := ConcordanceRequest{Text: text}

	// Now call Concordance with req.Text
	concordance := Concordance(req.Text)
	c.JSON(http.StatusOK, concordance)
}

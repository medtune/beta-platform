package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// SentimentAnalysis demo
func SentimentAnalysis(c *gin.Context) {
	c.Status(200)
	tmpl.DemoSentimentAnalysis.Execute(c.Writer, &data.Main{
		Version:   internal.VERSION,
		PageTitle: "Demo: Sentiment Analysis",
	})
}

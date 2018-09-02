package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// SentimentAnalysis demo
func SentimentAnalysis(c *gin.Context) {
	c.Status(200)
	tmpl.DemoSentimentAnalysis.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Sentiment Analysis",
	})
}

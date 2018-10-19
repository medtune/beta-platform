package platform

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/service/dashboard"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// Dashboard .
func Dashboard(c *gin.Context) {
	c.Status(200)
	versionB, _ := json.MarshalIndent(internal.GetVersion(), "", "    ")
	configB, _ := yaml.Marshal(dashboard.StartupConfig)
	capsules, _ := dashboard.GetCapsulesList()
	tmpl.Dashboard.Execute(c.Writer, &data.Dashboard{
		Title:    "Dashboard",
		Version:  string(versionB),
		Config:   string(configB),
		Capsules: capsules,
	})
}

package platform

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/service/dashboard"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Dashboard .
func Dashboard(c *gin.Context) {
	c.Status(200)
	versionB, _ := json.MarshalIndent(pkg.GetVersion(), "", "    ")
	configB, _ := yaml.Marshal(dashboard.StartupConfig)
	tmpl.Dashboard.Execute(c.Writer, &data.Dashboard{
		Title:   "Dashboard",
		Version: string(versionB),
		Config:  string(configB),
	})
}

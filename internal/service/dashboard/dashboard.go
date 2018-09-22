package dashboard

import (
	"time"

	"github.com/medtune/beta-platform/internal/config"
)

var (
	StartupConfig *config.StartupConfig
	StartupDate   time.Time
)

// GetCapsulesList .
func GetCapsulesList() ([]*config.ModelConfig, error) {
	cl := make([]*config.ModelConfig, 0, 20)
	cc := StartupConfig.Capsul
	ccc := StartupConfig.CustomCapsul

	if cc.Mnist != nil {
		cl = append(cl, cc.Mnist)
	}
	if cc.Inception != nil {
		cl = append(cl, cc.Inception)
	}
	if cc.MuraMNV2 != nil {
		cl = append(cl, cc.MuraMNV2)
	}
	if cc.MuraIRNV2 != nil {
		cl = append(cl, cc.MuraIRNV2)
	}
	if cc.ChexrayMNV2 != nil {
		cl = append(cl, cc.ChexrayMNV2)
	}
	if cc.ChexrayDN121 != nil {
		cl = append(cl, cc.ChexrayDN121)
	}
	if ccc.ChexrayDN121Cam != nil {
		cl = append(cl, ccc.ChexrayDN121Cam)
	}
	if ccc.ChexrayMNV2Cam != nil {
		cl = append(cl, ccc.ChexrayMNV2Cam)
	}
	if ccc.MuraIRNV2Cam != nil {
		cl = append(cl, ccc.MuraIRNV2Cam)
	}
	if ccc.MuraMNV2Cam != nil {
		cl = append(cl, ccc.MuraMNV2Cam)
	}
	if ccc.ChexrayPP != nil {
		cl = append(cl, ccc.ChexrayPP)
	}

	return cl, nil
}

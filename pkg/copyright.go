package pkg

import "strings"

var (
	Authors     string
	Owners      string
	LicenseURL  string
	LicenseType string
)

func init() {
	Owners = strings.Replace(strings.Replace(Owners, "/", ", ", -1), ".", " ", -1)
	Authors = strings.Replace(strings.Replace(Authors, "/", ", ", -1), ".", " ", -1)
}

// CopyrightInfo data
type CopyrightInfo struct {
	Authors     string `json:"Authors"`
	Owners      string `json:"Owners"`
	LicenseType string `json:"LicenseType"`
	LicenseURL  string `json:"LicenseURL"`
}

// GetCopyright .
func GetCopyright() *CopyrightInfo {
	return &CopyrightInfo{
		Authors:     Authors,
		Owners:      Owners,
		LicenseType: LicenseType,
		LicenseURL:  LicenseURL,
	}
}

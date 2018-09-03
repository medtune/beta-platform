package pkg

// GetCopyrights .
func GetCopyrights() *CopyrightInfo {
	return &CopyrightInfo{
		Authors:     Authors,
		Owners:      Owners,
		LicenseType: LicenseType,
		LicenseURL:  LicenseURL,
	}
}

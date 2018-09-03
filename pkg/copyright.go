package pkg

// GetCopyrights .
func GetCopyright() *CopyrightInfo {
	return &CopyrightInfo{
		Authors:     Authors,
		Owners:      Owners,
		LicenseType: LicenseType,
		LicenseURL:  LicenseURL,
	}
}

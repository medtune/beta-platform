package data

// Image struct
type Image struct {
	Filename    string
	Name        string
	Format      string
	Description string
}

// InceptionDemo data struct
type InceptionDemo struct {
	Main
	Samples []Image
}

// MuraV2Demo data struct
type MuraV2Demo struct {
	Main
	Samples []Image
}

// ChexrayV2Demo data struct
type ChexrayV2Demo struct {
	Main
	Samples []Image
}

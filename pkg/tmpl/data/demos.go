package data

// Image struct
type Image struct {
	URL         string
	Name        string
	Format      string
	Description string
}

// InceptionDemo data struct
type InceptionDemo struct {
	Main
	Samples []Image
}

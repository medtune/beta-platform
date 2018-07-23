package data

type Image struct {
	URL         string
	Name        string
	Format      string
	Description string
}

type InceptionDemo struct {
	Main
	Samples []Image
}

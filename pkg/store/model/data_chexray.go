package model

// PathologyAnalysisLevel .
type PathologyAnalysisLevel struct {
	Id           int64
	Name         string
	AnalysisCode string

	RedCells           string
	Hemoglobine        string
	Hematocrite        string
	VGM                string
	CCMH               string
	TCMH               string
	WhiteCells         string
	PNNeutrophiles     string
	PNBasophiles       string
	Lymphocytes        string
	Monocytes          string
	Plaquettes         string
	SedimentationSpeed string
	Creatinine         string
	GlycemieAjeun      string
	Triglycerides      string
	CholesterolTotal   string
	HDLCholesterol     string
	SGOTASAT           string
	SGPTALAT           string
	ProteinesUrine     string
	GlucoseUrine       string
	LeucocytesUrine    string
	HematiesUrines     string
}

// SpecAnalysisPool .
type SpecAnalysisPool struct {
	Id   int64
	Name string
	Unit string
	Min  int
	Max  int
}

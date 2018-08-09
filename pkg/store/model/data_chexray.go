package model

// PathologyAnalysisLevel .
type PathologyAnalysisLevel struct {
	Id                 int64
	Name               string
	RedCells           string
	Hemoglobin         string
	Hematocrit         string
	MCV                string
	MCHC               string
	MCH                string
	WhiteCells         string
	PNNeutrophils      string
	PNBasophils        string
	Lymphocytes        string
	Monocytes          string
	Platelet           string
	SedimentationSpeed string
	Creatinine         string
	GlucoseLevelAtFast string
	Triglycerides      string
	CholesterolLevel   string
	HDLCholesterol     string
	SGOTAST            string
	SGPTALT            string
	ProteinUrin        string
	GlucoseUrin        string
	LeukocyteUrin      string
	Hematuria          string
}

// SpecAnalysisPool .
type SpecAnalysisPool struct {
	Id   int64
	Name string
	Unit string
	Min  int64
	Max  int64
}

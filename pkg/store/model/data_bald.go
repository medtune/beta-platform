package model

// PathologyAnalysisLevel .
type PathologyAnalysisLevel struct {
	Id                 int64  `json:"id"`
	Name               string `xorm:"varchar(128) not null unique" json:"Name" valid:"required"`
	RedCells           string `xorm:"varchar(32) not null" json:"Red-Cells" valid:"required"`
	Hemoglobin         string `xorm:"varchar(32) not null" json:"Hemoglobin" valid:"required"`
	Hematocrit         string `xorm:"varchar(32) not null" json:"Hematocrit" valid:"required"`
	MCV                string `xorm:"varchar(32) not null" json:"VGM" valid:"required"`
	MCHC               string `xorm:"varchar(32) not null" json:"CCMH" valid:"required"`
	MCH                string `xorm:"varchar(32) not null" json:"TCMH" valid:"required"`
	WhiteCells         string `xorm:"varchar(32) not null" json:"White-Cells" valid:"required"`
	PNNeutrophils      string `xorm:"varchar(32) not null" json:"PN-Neutrophils" valid:"required"`
	PNBasophils        string `xorm:"varchar(32) not null" json:"PN-Basophils" valid:"required"`
	PNEosinophils      string `xorm:"varchar(32) not null" json:"PN-Eosinophils" valid:"required"`
	Lymphocytes        string `xorm:"varchar(32) not null" json:"Lymphocytes" valid:"required"`
	Monocytes          string `xorm:"varchar(32) not null" json:"Monocytes" valid:"required"`
	Platelet           string `xorm:"varchar(32) not null" json:"Platelets" valid:"required"`
	SedimentationSpeed string `xorm:"varchar(32) not null" json:"Sedimentation-Speed-1h" valid:"required"`
	Creatinine         string `xorm:"varchar(32) not null" json:"Creatinine" valid:"required"`
	GlucoseLevelAtFast string `xorm:"varchar(32) not null" json:"Glucose-Fast" valid:"required"`
	Triglycerides      string `xorm:"varchar(32) not null" json:"Triglycerides" valid:"required"`
	CholesterolLevel   string `xorm:"varchar(32) not null" json:"Cholesterol-Level" valid:"required"`
	HDLCholesterol     string `xorm:"varchar(32) not null" json:"HDL-Cholesterol" valid:"required"`
	SGOTAST            string `xorm:"varchar(32) not null" json:"SGOT-ASAT" valid:"required"`
	SGPTALT            string `xorm:"varchar(32) not null" json:"SGPT-ALAT" valid:"required"`
	GammaGT            string `xorm:"varchar(32) not null" json:"Gamma-GT" valid:"required"`
	ProteinUrin        string `xorm:"varchar(32) not null" json:"Proteins-Urin" valid:"required"`
	GlucoseUrin        string `xorm:"varchar(32) not null" json:"Glucose-Urin" valid:"required"`
	LeukocyteUrin      string `xorm:"varchar(32) not null" json:"Hematuria" valid:"required"`
	Hematuria          string `xorm:"varchar(32) not null" json:"Leukocyte-Urin" valid:"required"`
}

// SpecAnalysisPool .
type SpecAnalysisPool struct {
	Id   int64  `json:"id"`
	Name string `xorm:"varchar(128) not null unique" json:"name" valid:"required"`
	Unit string `xorm:"varchar(32) not null" json:"unit" valid:"required"`
	Min  string `xorm:"varchar(32) not null" json:"min"`
	Max  string `xorm:"varchar(32) not null" json:"max"`
}

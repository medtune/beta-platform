package model

// PathologyAnalysisLevel .
type PathologyAnalysisLevel struct {
	Id   int64
	Name string `xorm:"varchar(128) not null unique" valid:"required"`

	RedCells           string `xorm:"varchar(32) not null" valid:"required"`
	Hemoglobin         string `xorm:"varchar(32) not null" valid:"required"`
	Hematocrit         string `xorm:"varchar(32) not null" valid:"required"`
	MCV                string `xorm:"varchar(32) not null" valid:"required"`
	MCHC               string `xorm:"varchar(32) not null" valid:"required"`
	MCH                string `xorm:"varchar(32) not null" valid:"required"`
	WhiteCells         string `xorm:"varchar(32) not null" valid:"required"`
	PNNeutrophils      string `xorm:"varchar(32) not null" valid:"required"`
	PNBasophils        string `xorm:"varchar(32) not null" valid:"required"`
	Lymphocytes        string `xorm:"varchar(32) not null" valid:"required"`
	Monocytes          string `xorm:"varchar(32) not null" valid:"required"`
	Platelet           string `xorm:"varchar(32) not null" valid:"required"`
	SedimentationSpeed string `xorm:"varchar(32) not null" valid:"required"`
	Creatinine         string `xorm:"varchar(32) not null" valid:"required"`
	GlucoseLevelAtFast string `xorm:"varchar(32) not null" valid:"required"`
	Triglycerides      string `xorm:"varchar(32) not null" valid:"required"`
	CholesterolLevel   string `xorm:"varchar(32) not null" valid:"required"`
	HDLCholesterol     string `xorm:"varchar(32) not null" valid:"required"`
	SGOTAST            string `xorm:"varchar(32) not null" valid:"required"`
	SGPTALT            string `xorm:"varchar(32) not null" valid:"required"`
	GammaGT            string `xorm:"varchar(32) not null" valid:"required"`
	ProteinUrin        string `xorm:"varchar(32) not null" valid:"required"`
	GlucoseUrin        string `xorm:"varchar(32) not null" valid:"required"`
	LeukocyteUrin      string `xorm:"varchar(32) not null" valid:"required"`
	Hematuria          string `xorm:"varchar(32) not null" valid:"required"`
}

// SpecAnalysisPool .
type SpecAnalysisPool struct {
	Id   int64
	Name string  `xorm:"varchar(128) not null unique" valid:"required"`
	Unit string  `xorm:"varchar(32) not null" valid:"required"`
	Min  float64 `xorm:"default -1"`
	Max  float64 `xorm:"default -1"`
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name       string
	Department string

	Orders []Order `gorm:"foreignKey:UserRefer"`
}

type Order struct {
	gorm.Model

	Order_name string `gorm:"unique"`
	UserRefer  uint

	Samples       []Sample      `gorm:"foreignKey:OrderRefer"`
	AnalysisSuite AnalysisSuite `gorm:"foreignKey:OrderRefer"`
	AnalysisRun   AnalysisRun   `gorm:"foreignKey:OrderRefer"`
	Anslysis      Analysis      `gorm:"foreignKey:OrderRefer"`
}

type Sample struct {
	gorm.Model

	Sample_name  string `gorm:"unique"`
	Sample_type  string
	Sample_state string
	OrderRefer   uint

	SampleFiles []SampleFile `gorm:"foreignKey:SampleRefer"`
}

type SampleFile struct {
	gorm.Model

	File_name     string `gorm:"unique"`
	File_size     int
	File_location string
	File_state    string
	SampleRefer   uint
}

type Analysis struct {
	gorm.Model

	Analysis_name string `gorm:"unique"`
	Analysis_type string

	AnalysisSuiteRefer uint
	OrderRefer         uint
}

type AnalysisSuite struct {
	gorm.Model

	Suite_name string `gorm:"unique"`
	OrderRefer uint

	Analyses []Analysis `gorm:"foreignKey:AnalysisSuiteRefer"`
}

type AnalysisRun struct {
	gorm.Model

	Analysis_run_id     uint `gorm:"autoIncrement"`
	Analysis_run_status string
	Analysis_run_type   string

	OrderRefer uint

	AnalysisRunOutput []AnalysisRunOutput `gorm:"foreignKey:AnalysisRunRefer"`
}

type AnalysisRunOutput struct {
	gorm.Model

	Output_file_name     string `gorm:"unique"`
	Output_file_size     int
	Output_file_location string
	Output_file_state    string

	AnalysisRunRefer uint
}

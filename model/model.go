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

	Order_name         string `gorm:"unique"`
	UserRefer          uint
	SampleRefer        uint
	AnalysisSuiteRefer uint
	AnalysisRunRefer   uint
	AnslysisRefer      uint

	Sample        Sample        `gorm:"foreignKey:SampleRefer"`
	AnalysisSuite AnalysisSuite `gorm:"foreignKey:AnalysisSuiteRefer"`
	AnalysisRun   AnalysisRun   `gorm:"foreignKey:AnalysisRunRefer"`
	Anslysis      Analysis      `gorm:"foreignKey:AnslysisRefer"`
}

type Sample struct {
	gorm.Model

	Sample_name  string `gorm:"unique"`
	Sample_type  string
	Sample_state string

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
}

type AnalysisSuite struct {
	gorm.Model

	Suite_name string `gorm:"unique"`

	Analyses []Analysis `gorm:"foreignKey:AnalysisSuiteRefer"`
}

type AnalysisRun struct {
	gorm.Model

	Analysis_run_id     uint `gorm:"autoIncrement"`
	Analysis_run_status string
	Analysis_run_type   string

	AnalysisRunOutput []AnalysisRunOutput `gorm:"foreignKey:AnalysisRunRefer"`
}

type AnalysisRunOutput struct {
	gorm.Model

	Output_file_name     string `gorm:"unique"`
	output_file_size     int
	output_file_location string
	output_file_state    string

	AnalysisRunRefer uint
}

package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	name       string
	department string
	username   string `gorm:"unique"`

	Orders []Order `gorm:"foreignKey:UserRefer"`
}

type Order struct {
	gorm.Model

	order_name string `gorm:"unique"`
	UserRefer  uint

	Samples       []Sample      `gorm:"foreignKey:OrderRefer"`
	AnalysisSuite AnalysisSuite `gorm:"foreignKey:OrderRefer"`
	AnalysisRun   AnalysisRun   `gorm:"foreignKey:OrderRefer"`
}

type Sample struct {
	gorm.Model

	sample_name  string `gorm:"unique"`
	sample_type  string
	sample_state string
	OrderRefer   uint

	SampleFiles []SampleFile `gorm:"foreignKey:SampleRefer"`
}

type SampleFile struct {
	gorm.Model

	file_name     string `gorm:"unique"`
	file_size     int
	file_location string
	file_state    string
	SampleRefer   uint
}

type Analysis struct {
	gorm.Model

	analysis_name string `gorm:"unique"`
	analysis_type string

	AnalysisSuiteRefer uint
}

type AnalysisSuite struct {
	gorm.Model

	suite_name string `gorm:"unique"`
	OrderRefer uint

	Analyses []Analysis `gorm:"foreignKey:AnalysisSuiteRefer"`
}

type AnalysisRun struct {
	gorm.Model

	analysis_run_id     uint `gorm:"autoIncrement"`
	analysis_run_status string
	analysis_run_type   string

	OrderRefer uint

	AnalysisRunOutput []AnalysisRunOutput `gorm:"foreignKey:AnalysisRunRefer"`
}

type AnalysisRunOutput struct {
	gorm.Model

	output_file_name     string `gorm:"unique"`
	output_file_size     int
	output_file_location string
	output_file_state    string

	AnalysisRunRefer uint
}

func main() {
	dsn := "host=localhost user=fileadmin password=test1234 dbname=fileflipperdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		db.AutoMigrate(
			&User{},
			&Order{},
			&Sample{},
			&SampleFile{},
			&Analysis{},
			&AnalysisSuite{},
			&AnalysisRun{},
			&AnalysisRunOutput{},
		)
		print("connected to database", db)
	}
}

package main

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID         uint `gorm:"primaryKey;default:auto_random()"`
	name       string
	department string
	username   string
	createAt   time.Time
	updateAt   time.Time

	Orders []Order `gorm:"foreignKey:UserRefer"`
}

type Order struct {
	ID         uint `gorm:"primaryKey;default:auto_random()"`
	order_name string
	createAt   time.Time
	updateAt   time.Time

	Samples       []Sample      `gorm:"foreignKey:OrderRefer"`
	AnalysisSuite AnalysisSuite `gorm:"foreignKey:OrderRefer"`
	AnalysisRun   AnalysisRun   `gorm:"foreignKey:OrderRefer"`
}

type Sample struct {
	ID          uint `gorm:"primaryKey;default:auto_random()"`
	sample_name string
	createAt    time.Time
	updateAt    time.Time

	SampleFiles []SampleFile `gorm:"foreignKey:SampleRefer"`
}

type SampleFile struct {
	ID            uint `gorm:"primaryKey;default:auto_random()"`
	file_name     string
	file_size     int
	file_location string
	file_state    string
	createAt      time.Time
	updateAt      time.Time
}

type Analysis struct {
	ID            uint `gorm:"primaryKey;default:auto_random()"`
	analysis_name string
	analysis_type string
	createAt      time.Time
	updateAt      time.Time
}

type AnalysisSuite struct {
	ID         uint `gorm:"primaryKey;default:auto_random()"`
	suite_name string
	createAt   time.Time
	updateAt   time.Time

	analysis []Analysis `gorm:"foreignKey:SuiteRefer"`
}

type AnalysisRun struct {
	ID                  uint `gorm:"primaryKey;default:auto_random()"`
	analysis_run_id     string
	analysis_run_status string

	analysis_run_type time.Time
	createAt          time.Time
	updateAt          time.Time

	AnalysisRunOutput []AnalysisRunOutput `gorm:"foreignKey:RunRefer"`
}

type AnalysisRunOutput struct {
	ID                   uint `gorm:"primaryKey;default:auto_random()"`
	output_file_name     string
	output_file_size     int
	output_file_location string
	output_file_state    string
	createAt             time.Time
	updateAt             time.Time
}

func main() {
	dsn := "host=localhost user=fileadmin password=test1234 dbname=fileflipperdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		print("connected to database", db)
	}
}

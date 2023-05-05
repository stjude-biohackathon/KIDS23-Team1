package db

import "time"

const TIMELAYOUT = "2006-11-30"

type AnalysisSuite struct {
	ID        int    `json:"id"`
	SuiteName string `json:"suite_name"`
}

type AnalysisType struct {
	ID               int    `json:"id"`
	AnalysisTypeName string `json:"anls_type_name"`
	AnalysisSuiteID  int    `json:"anls_suite_id"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Order struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	AnalysisSuiteID int       `json:"anls_suite_id"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

type OrderSample struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	SampleID int `json:"sample_id"`
}

type AnalysisRun struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	AnalysisId    int       `json:"anls_id"`
	OrderSampleID int       `json:"order_sample_id"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type AnalysisRunOutputs struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	AnalysisRunID    int    `json:"anls_run_id"`
	FilesystemPath   string `json:"filesystem_path"`
	LocationEncoding string `json:"location_encoding"`
	CreatedAt        string `json:"created_at"`
}

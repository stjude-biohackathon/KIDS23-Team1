package db

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
